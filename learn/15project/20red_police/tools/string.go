package tools

import "github.com/agext/uuid"

func UUID() string {
	return uuid.New().String()
}
