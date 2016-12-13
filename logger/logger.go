package logger

import (
	"fmt"
	"time"
)

var debug = true

// Logger function
func Logger(funcName string, msg string) {
	if debug {
		fmt.Printf("%s %s() - %s\n\n", time.Now().Format("2006/01/02 15:04:05"), funcName, msg)
	}
}
