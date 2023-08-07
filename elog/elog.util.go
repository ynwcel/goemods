package elog

import (
	"fmt"
	"runtime"
	"strings"
)

func kv2map(kv ...any) Map {
	if len(kv)%2 != 0 {
		kv = append(kv, miss_msg)
	}
	mapVal := make(Map)
	for i := 0; i < len(kv); i += 2 {
		mapVal[fmt.Sprintf("%v", kv[i])] = kv[i+1]
	}
	return mapVal
}

func getFileLine() (string, int) {
	find_file, find_line := "", -1
	skip := 2
	for {
		if _, file, line, ok := runtime.Caller(skip); !ok {
			break
		} else if !strings.Contains(file, call_skip) {
			find_file, find_line = file, line
			break
		}
		skip += 1
		if skip >= 10 {
			break
		}
	}
	return find_file, find_line
}
