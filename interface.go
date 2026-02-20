package emi_core

import (
	"context"

	"github.com/aK1r4z/emi-core/api"
	milky "github.com/aK1r4z/emi-core/types"
)

type Logger interface {
	Tracef(format string, args ...any)
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)

	Trace(args ...any)
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)
}

type APIClient interface {
	SendPrivateMessage(ctx context.Context, request api.SendPrivateMessageRequest) (*api.SendPrivateMessageResponse, error)
	SendGroupMessage(ctx context.Context, request api.SendGroupMessageRequest) (*api.SendPrivateMessageResponse, error)
	// [TODO] 包装更多 API 方法
}

type Bot interface {
	Open(context.Context) error
	Close() error

	Wait()

	APIClient

	Handle(EventHandler)

	Logger() Logger
}

type EventSource interface {
	Open() (chan milky.RawEvent, error)
	Close() error
}

type EventHandler interface {
	milky.EventDescriber
	Handle(Bot, milky.Event, milky.RawEvent)
}
