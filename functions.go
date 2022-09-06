package xcommon

import (
	"github.com/spf13/cast"
	"log"
	"runtime/debug"
)

func Recover() {
	if rec := recover(); rec != nil {
		if err, ok := rec.(error); ok {
			log.Printf("PanicRecover Unhandled error: %v\n stack:%v \n", err.Error(), cast.ToString(debug.Stack()))
		} else {
			log.Printf("PanicRecover Panic: %v\n stack:%v \n", rec, cast.ToString(debug.Stack()))
		}
	}
}
