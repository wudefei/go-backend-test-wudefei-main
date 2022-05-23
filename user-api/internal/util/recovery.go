package util

import (
	"fmt"
	"runtime/debug"
	"time"
)

// pretty panic
func PrettyPanic(logId string, msg string, r interface{}) {
	var (
		stacks = debug.Stack()
		t      = time.Now().Format("2006-01-02 15:04:05")
	)
	errorStr := fmt.Sprintf("\n%v %s\n[Recovery]\npanic:%v\n\nstacks:\n%s", t, msg, r, string(stacks))
	fmt.Println(logId, errorStr)
}
