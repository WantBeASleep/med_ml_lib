package dbus

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"google.golang.org/protobuf/proto"
)

// SendOptions вызова переопределят SendOptions продюсера
type Producer[T proto.Message] interface {
	Send(ctx context.Context, msg T) error
	Close() error
}

type producer[T proto.Message] struct {
	producer sarama.SyncProducer

	topic string

	// будут вызваны при отправке сообщения
	middlewares []producerMiddlewars[T]
}

type producerMiddlewars[T proto.Message] func(
	ctx context.Context,
	topic string,
	msg T,
	cfg *OptionalSendCfg,
) error

func NewProducer[T proto.Message](
	saramaProducer sarama.SyncProducer,
	topic string,
	middlewares ...producerMiddlewars[T],
) Producer[T] {
	return &producer[T]{
		producer:    saramaProducer,
		topic:       topic,
		middlewares: middlewares,
	}
}

// конфиг для отправки сообщения
type OptionalSendCfg struct {
	// ключ для партиционирования
	Key *string
	// точная партиция
	Partition *int
	Headers   map[string]string
}

func (cfg *OptionalSendCfg) Apply(msg *sarama.ProducerMessage) {
	if cfg.Partition != nil {
		msg.Partition = int32(*cfg.Partition)
	}
	if cfg.Key != nil {
		msg.Key = sarama.StringEncoder(*cfg.Key)
	}
	for k, v := range cfg.Headers {
		msg.Headers = append(msg.Headers, sarama.RecordHeader{
			Key:   []byte(k),
			Value: []byte(v),
		})
	}
}

func (p *producer[T]) Send(ctx context.Context, msg T) error {
	payload, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal message: %w", err)
	}

	saramaMsg := &sarama.ProducerMessage{
		Topic: p.topic,
		Key:   sarama.StringEncoder(""), // default value
		Value: sarama.ByteEncoder(payload),
	}

	cfg := &OptionalSendCfg{}
	for _, mdlwr := range p.middlewares {
		if err := mdlwr(ctx, p.topic, msg, cfg); err != nil {
			return fmt.Errorf("run middleware: %w", err)
		}
	}

	for k, v := range cfg.Headers {
		saramaMsg.Headers = append(saramaMsg.Headers, sarama.RecordHeader{
			Key:   []byte(k),
			Value: []byte(v),
		})
	}

	cfg.Apply(saramaMsg)

	_, _, err = p.producer.SendMessage(saramaMsg)
	if err != nil {
		return fmt.Errorf("send message: %w", err)
	}

	return nil
}

func (p *producer[T]) Close() error {
	return p.producer.Close()
}
