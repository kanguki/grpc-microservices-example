package log

import (
	"fmt"
	l "log"
)

func Log(formatOrValue interface{}, values ...interface{}) {
	//TODO: send log to mq
	l.Printf(fmt.Sprintf("<DIV>: %v", formatOrValue), values...)
}
