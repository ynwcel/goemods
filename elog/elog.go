package elog

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/gogf/gf/v2/os/gctx"
)

const (
	miss_msg  = "(!MISS)"
	call_skip = "goecore/elog"
)

type (
	Map    = map[string]any
	logger struct {
		ctxKeys  []string
		logpath  string
		filename string
		openFile *os.File
		sync.RWMutex
		addFileLine bool
	}
	Option func(*logger)
)

func OptSetLogPath(logpath string) Option {
	return func(l *logger) {
		l.logpath = logpath
	}
}

func OptAddFileLine(addFileLine bool) Option {
	return func(l *logger) {
		l.addFileLine = addFileLine
	}
}

func New(options ...Option) *logger {
	l := &logger{}
	for _, opt := range options {
		opt(l)
	}
	return l
}

func (l *logger) DebugKv(ctx context.Context, kv ...any) error {
	return l.DebugMap(ctx, kv2map(kv...))
}

func (l *logger) InfoKv(ctx context.Context, kv ...any) error {
	return l.InfoMap(ctx, kv2map(kv...))
}

func (l *logger) WarnKv(ctx context.Context, kv ...any) error {
	return l.WarnMap(ctx, kv2map(kv...))
}

func (l *logger) ErrorKv(ctx context.Context, kv ...any) error {
	return l.ErrorMap(ctx, kv2map(kv...))
}

func (l *logger) DebugMap(ctx context.Context, mapVal Map) error {
	return l.record(ctx, "Debug", mapVal)
}

func (l *logger) InfoMap(ctx context.Context, mapVal Map) error {
	return l.record(ctx, "Info", mapVal)
}

func (l *logger) WarnMap(ctx context.Context, mapVal Map) error {
	return l.record(ctx, "Warn", mapVal)
}

func (l *logger) ErrorMap(ctx context.Context, mapVal Map) error {
	return l.record(ctx, "Error", mapVal)
}

func (l *logger) AddCtxKey(keys ...string) {
	l.Lock()
	defer l.Unlock()
	l.ctxKeys = append(l.ctxKeys, keys...)
}

func (l *logger) record(ctx context.Context, level string, mapVal Map) error {
	w, err := l.getWriter()
	if err != nil {
		return err
	}
	l.RLock()
	mapVal["time"] = time.Now().Format("2006-01-02 15:04:05")
	mapVal["level"] = level
	mapVal["trace_id"] = gctx.CtxId(ctx)
	if len(l.ctxKeys) > 0 {
		for _, key := range l.ctxKeys {
			mapVal[key] = ctx.Value(key)
		}
	}
	if l.addFileLine {
		sfile, sline := getFileLine()
		mapVal["file"] = sfile
		mapVal["line"] = sline
	}
	l.RUnlock()
	if jsonBytes, err := json.Marshal(mapVal); err != nil {
		return err
	} else {
		_, err = fmt.Fprintln(w, string(jsonBytes))
		return err
	}
}

func (l *logger) getWriter() (io.Writer, error) {
	if len(l.logpath) <= 0 {
		return os.Stdout, nil
	}
	l.Lock()
	defer l.Unlock()
	cur_filename := time.Now().Format("2006-01-02")
	if cur_filename != l.filename {
		if l.openFile != nil {
			l.openFile.Close()
		}
		l.filename = cur_filename
		log_file := fmt.Sprintf("%s/%s.log", l.logpath, l.filename)
		if of, err := os.OpenFile(log_file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err != nil {
			return nil, err
		} else {
			l.openFile = of
		}
	}
	return l.openFile, nil
}
