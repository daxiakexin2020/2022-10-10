package tools

import "time"

func FormatTimeToString(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

func NowTimeFormatTimeToString() string {
	return FormatTimeToString(time.Now())
}
