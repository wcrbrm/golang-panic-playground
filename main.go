package main

import (
	"fmt"
	"runtime"
)

func somethingThatWillBreak() {
	arr := []string{"one", "two"}
	_ = arr[10]
}

func wrapper() {
	// attempt to recover from panic
	if r := recover(); r != nil {
		buf := make([]byte, 1<<16)
		runtime.Stack(buf, true)

		fmt.Printf("panic: %+v", r)
		fmt.Printf("stacktrace: %s", buf)
	}
}

func main() {
	defer wrapper()

	somethingThatWillBreak()
}
