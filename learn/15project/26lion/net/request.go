package net

import (
	"26lion/iface"
	"sync"
)

const (
	PRE_HANDLE iface.HandleStep = iota
	HANDLE
	POST_HANDLE

	HANDLE_OVER
)

type Request struct {
	conn     iface.IConnection
	router   iface.IRouter
	steps    iface.HandleStep
	stepLock *sync.RWMutex
	needNext bool
}
