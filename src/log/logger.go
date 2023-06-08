package log

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

func Setup() {
	log.SetFlags(log.Ldate | log.Ltime)
}

func Infof(msg string, args ...any) {
	txt := fmt.Sprintf(msg, args...)
	fmt.Printf("\x1b[35m%s\x1b[0m\n", txt)
}

func Errorf(err error, msg string, args ...any) {
	txt := fmt.Sprintf(msg, args...)
	log.Printf("\x1b[33m[ERROR] %s %s %s\x1b[0m", getFileLine(), txt, err.Error())
}

func getFileLine() string {
	var ret string
	if _, file, line, ok := runtime.Caller(2); ok {
		parts := strings.Split(file, "/")
		length := len(parts)
		ret = fmt.Sprintf("%s/%s:%d", parts[length-2], parts[length-1], line)
	}
	return ret
}
