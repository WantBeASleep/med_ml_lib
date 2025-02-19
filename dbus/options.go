package dbus

import (
	"google.golang.org/protobuf/proto"
)

// ---- Subscriber options ----
type SubscriberOptions[T proto.Message] func(*handler[T])

func WithInitFunc[T proto.Message](initFuncs ...initFunc) SubscriberOptions[T] {
	return func(h *handler[T]) {
		h.initFuncs = append(h.initFuncs, initFuncs...)
	}
}

func WithSubscriberMiddlewars[T proto.Message](preConsumeFuncs ...subscriberMiddlewars[T]) SubscriberOptions[T] {
	return func(h *handler[T]) {
		h.preConsumeFuncs = append(h.preConsumeFuncs, preConsumeFuncs...)
	}
}
