package construct

import "time"

type exTime struct {
	time int64
}

func (ex *exTime) GetExTime() int64 {
	return ex.time
}

func (ex *exTime) SetExTime(t time.Duration) bool {
	ex.time = time.Now().Add(time.Second * t).Unix()
	return true
}

func NewExTime(t time.Duration) *exTime {
	return &exTime{time: time.Now().Add(time.Second * t).Unix()}
}
