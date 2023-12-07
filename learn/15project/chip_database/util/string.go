package util

import "github.com/rs/xid"

func UniqId() string {
	return xid.New().String()
}
