package tools

import (
	"fmt"
	"os"
)

func Error(msg string) {
	fmt.Fprintf(os.Stderr, "\033[31m%s\033[m\n", msg)
}

func Errorf(format string, msg interface{}) {
	fmt.Fprintf(os.Stderr, "\033[31m%s\033[m\n", fmt.Sprintf(format, msg))
}

func Warn(msg string) {
	fmt.Fprintf(os.Stderr, "\033[33m%s\033[m\n", msg)
}

func Warnf(format string, msg interface{}) {
	fmt.Fprintf(os.Stderr, "\033[33m%s\033[m\n", fmt.Sprintf(format, msg))
}

func Success(msg string) {
	fmt.Fprintf(os.Stderr, "\033[36m%s\033[m\n", msg)
}

func Successf(format string, msg interface{}) {
	fmt.Fprintf(os.Stderr, "\033[36m%s\033[m\n", fmt.Sprintf(format, msg))
}
