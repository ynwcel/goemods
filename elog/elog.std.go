package elog

import "context"

var (
	std_log = New()
)

func DebugKv(ctx context.Context, kv ...any) error {
	return std_log.DebugKv(ctx, kv...)
}

func InfoKv(ctx context.Context, kv ...any) error {
	return std_log.InfoKv(ctx, kv...)
}

func ErrorKv(ctx context.Context, kv ...any) error {
	return std_log.ErrorKv(ctx, kv...)
}

func WarnKv(ctx context.Context, kv ...any) error {
	return std_log.WarnKv(ctx, kv...)
}

func DebugMap(ctx context.Context, mapVal Map) error {
	return std_log.DebugMap(ctx, mapVal)
}

func InfoMap(ctx context.Context, mapVal Map) error {
	return std_log.InfoMap(ctx, mapVal)

}

func ErrorMap(ctx context.Context, mapVal Map) error {
	return std_log.ErrorMap(ctx, mapVal)

}

func WarnMap(ctx context.Context, mapVal Map) error {
	return std_log.WarnMap(ctx, mapVal)
}

func AddCtxKey(keys ...string) {
	std_log.AddCtxKey(keys...)
}

func SetStdLogger(log *logger) {
	std_log = log
}
