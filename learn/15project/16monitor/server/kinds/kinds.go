package kinds

import "time"

type Base struct {
	percent       int
	timeInterval  time.Duration
	lastAlarmTime int64
}
