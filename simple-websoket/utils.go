package simple_websoket

import (
	"fmt"
	"runtime/debug"
)

func ErrRecover(args ...interface{}) {
	if err := recover(); err != nil {
		if len(args) > 0 {
			s := fmt.Sprintln(args...)
			fmt.Println(err, s, string(debug.Stack()))
		} else {
			fmt.Println(err, string(debug.Stack()))
		}
	}
}
