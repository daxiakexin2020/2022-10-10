package error

import (
	"fmt"
)

var (
	TASk_YET_EXISTS   = "%s: task yet exists"
	TASk_NOT_EXISTS   = "%s: task not exists"
	TASK_LENGTH_ERROR = "%s: task length error"
)

func NewError(format string, a ...any) error {
	return fmt.Errorf(format, a)
}

func TaskYetExistsErr(s string) error {
	return NewError(TASk_YET_EXISTS, s)
}

func TaskNotExistsErr(s string) error {
	return NewError(TASk_NOT_EXISTS, s)
}

func TaskLengthError(s string) error {
	return NewError(TASK_LENGTH_ERROR, s)
}
