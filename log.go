package main

import (
	"fmt"
	"os"
)

func fatal(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Printf("\n")
	os.Exit(1)
}

func info(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Printf("\n")
}
