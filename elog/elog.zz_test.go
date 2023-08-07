package elog

import (
	"context"
	"testing"
	"time"

	"github.com/gogf/gf/v2/os/gctx"
)

func test_logger(t *testing.T, log *logger, ctx context.Context) {
	t.Parallel()
	kvfuns := []func(context.Context, ...any) error{
		log.DebugKv,
		log.InfoKv,
		log.WarnKv,
		log.ErrorKv,
	}
	mapfuncs := []func(context.Context, Map) error{
		log.DebugMap,
		log.InfoMap,
		log.WarnMap,
		log.ErrorMap,
	}
	kvargs := [][]any{
		[]any{"key1", "val1", "key2"},
		[]any{"key1", "val1", "key2", "val2", "num", 1101},
	}
	mapargs := []Map{
		Map{"key1": "val1", "key2": ""},
		Map{"key1": "val1", "key2": "val2", "num": 1101},
	}
	for _, f := range kvfuns {
		for _, arg := range kvargs {
			if err := f(ctx, arg...); err != nil {
				t.Error(err)
			}
		}
	}
	for _, f := range mapfuncs {
		for _, arg := range mapargs {
			if err := f(ctx, arg); err != nil {
				t.Error(err)
			}
		}
	}
}

func TestStd(t *testing.T) {
	test_logger(t, std_log, gctx.New())
}

func TestStdCtxKey(t *testing.T) {
	AddCtxKey("client_ip", "request_id")
	ctx := gctx.New()
	ctx = context.WithValue(ctx, "client_ip", "127.0.0.1")
	test_logger(t, std_log, ctx)
}

func TestFileLog(t *testing.T) {
	logger := New(OptSetLogPath("./"), OptAddFileLine(true))
	test_logger(t, logger, gctx.New())
}

func TestFileCtxKey(t *testing.T) {
	logger := New(OptSetLogPath("./"), OptAddFileLine(true))
	logger.AddCtxKey("client_ip", "request_id")
	ctx := gctx.New()
	ctx = context.WithValue(ctx, "client_ip", "127.0.0.1")
	ctx = context.WithValue(ctx, "request_id", time.Now().Unix())
	test_logger(t, logger, ctx)
}
