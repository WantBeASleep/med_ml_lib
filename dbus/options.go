package dbus

import (
	"google.golang.org/protobuf/proto"
)

// ---- Subscriber options ----
type SubscriberOptions[T proto.Message] func(*handler[T])

func WithInitMiddlewares[T proto.Message](initMiddlewares ...subscriberInitMiddlewares) SubscriberOptions[T] {
	return func(h *handler[T]) {
		h.initMiddlewars = append(h.initMiddlewars, initMiddlewares...)
	}
}

func WithSubscriberTypeMiddlewares[T proto.Message](typeMiddlewares ...subscriberTypeMiddlewares[T]) SubscriberOptions[T] {
	return func(h *handler[T]) {
		h.typeMiddlewars = append(h.typeMiddlewars, typeMiddlewares...)
	}
}

func WithSubscriberMiddlewares[T proto.Message](middlewares ...subscriberMiddlewares) SubscriberOptions[T] {
	return func(h *handler[T]) {
		h.middlewares = append(h.middlewares, middlewares...)
	}
}

// ---- Producer options ----
type ProducerOptions[T proto.Message] func(*producer[T])

func WithProducerMiddlewares[T proto.Message](middlewares ...producerMiddlewares) ProducerOptions[T] {
	return func(p *producer[T]) {
		p.middlewares = append(p.middlewares, middlewares...)
	}
}

func WithProducerTypeMiddlewares[T proto.Message](typeMiddlewares ...producerTypeMiddlewares[T]) ProducerOptions[T] {
	return func(p *producer[T]) {
		p.typeMiddlewares = append(p.typeMiddlewares, typeMiddlewares...)
	}
}
