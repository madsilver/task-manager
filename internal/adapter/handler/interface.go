package handler

import (
	"context"
	"time"
)

type ContextHTTP interface {
	Param(key string) string
	QueryParam(key string) string
	Path() string
	AddParam(key, value string)
	Query(key string) string
	JSON(code int, i any) error
	Bind(i interface{}) error
	Data(code int, contentType string, data []byte) error
	Abort(code int)
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key any) any
}

type Logger interface {
	Info(ctx context.Context, msg string)
	Warn(ctx context.Context, msg string)
	Error(ctx context.Context, msg string, err error)
	Panic(ctx context.Context, msg string)
}
