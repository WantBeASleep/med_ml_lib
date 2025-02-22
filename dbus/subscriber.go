package dbus

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/IBM/sarama"
	"google.golang.org/protobuf/proto"
)

type Subscriber interface {
	Start(ctx context.Context) error
	Close() error
}

type subscriber[T proto.Message] struct {
	topic   string
	hosts   []string
	groupID string

	handler *handler[T]
}

type SubscriberInitConnectionStats struct {
	Topic         string
	Partition     int
	InitialOffset int
}

type EventStats[T proto.Message] struct {
	Headers   map[string]string
	Key       string
	Value     T
	Topic     string
	Partition int
	Offset    int
}

type subscriberInitMiddlewares func(ctx context.Context, stats SubscriberInitConnectionStats) error
type subscriberMiddlewares func(ctx context.Context, stats EventStats[proto.Message]) (context.Context, error)
type subscriberTypeMiddlewares[T proto.Message] func(ctx context.Context, stats EventStats[T]) (context.Context, error)

type handler[T proto.Message] struct {
	// десереализованное сообщение
	consumer Consumer[T]

	// будут вызваны при инициализации подключения
	initMiddlewars []subscriberInitMiddlewares

	// будут вызваны при получении сообщения
	middlewares    []subscriberMiddlewares
	typeMiddlewars []subscriberTypeMiddlewares[T]
}

type Consumer[T proto.Message] interface {
	Consume(ctx context.Context, event T) error
}

func NewGroupSubscriber[T proto.Message](
	topic string,
	hosts []string,
	groupID string,
	consumer Consumer[T],
	options ...SubscriberOptions[T],
) Subscriber {
	h := &handler[T]{consumer: consumer}
	for _, option := range options {
		option(h)
	}

	return &subscriber[T]{
		topic:   topic,
		hosts:   hosts,
		groupID: groupID,
		handler: h,
	}
}

// Sarama interface
func (*handler[T]) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (*handler[T]) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (h *handler[T]) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) (err error) {
	defer func() {
		if r := recover(); r != nil {
			slog.ErrorContext(session.Context(), "panic in consumer group", slog.Any("error", r))
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	subStats := SubscriberInitConnectionStats{
		Topic:         claim.Topic(),
		Partition:     int(claim.Partition()),
		InitialOffset: int(claim.InitialOffset()),
	}

	for _, mdlwr := range h.initMiddlewars {
		if err := mdlwr(session.Context(), subStats); err != nil {
			return fmt.Errorf("run init middlewares: %w", err)
		}
	}

	for msg := range claim.Messages() {
		// T - proto.Message, дженерик изначально будет от указателя
		var event T
		if err := proto.Unmarshal(msg.Value, event); err != nil {
			return fmt.Errorf("unmarshal event to %T: %w", event, err)
		}

		eventStats := EventStats[T]{
			Topic:     msg.Topic,
			Key:       string(msg.Key),
			Value:     event,
			Partition: int(msg.Partition),
			Offset:    int(msg.Offset),
		}
		for _, v := range msg.Headers {
			eventStats.Headers[string(v.Key)] = string(v.Value)
		}

		ctx := session.Context()
		for _, mdlwr := range h.typeMiddlewars {
			var err error
			ctx, err = mdlwr(ctx, eventStats)
			if err != nil {
				return fmt.Errorf("run type middlewares: %w", err)
			}
		}

		for _, mdlwr := range h.middlewares {
			var err error
			ctx, err = mdlwr(ctx, EventStats[proto.Message]{
				Topic:     eventStats.Topic,
				Key:       eventStats.Key,
				Value:     eventStats.Value,
				Partition: eventStats.Partition,
				Offset:    eventStats.Offset,
				Headers:   eventStats.Headers,
			})
			if err != nil {
				return fmt.Errorf("run middlewares: %w", err)
			}
		}

		if err := h.consumer.Consume(ctx, event); err != nil {
			return fmt.Errorf("consume event: %w", err)
		}

		session.MarkMessage(msg, "")
		session.Commit()
	}

	return nil
}

func (s *subscriber[T]) Start(ctx context.Context) error {
	consumer, err := sarama.NewConsumerGroup(s.hosts, s.groupID, nil)
	if err != nil {
		return fmt.Errorf("create new group: %w", err)
	}

	for {
		if err := consumer.Consume(ctx, []string{s.topic}, s.handler); err != nil {
			return fmt.Errorf("consume topic: %w", err)
		}
	}
}

func (s *subscriber[T]) Close() error { return nil }
