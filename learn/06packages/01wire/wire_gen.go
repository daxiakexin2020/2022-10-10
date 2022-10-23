// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"01wire/service"
)

// Injectors from wire.go:

func InitializeC(msg string) service.C {
	a := service.A{
		Msg: msg,
	}
	b := service.B{
		AY: a,
	}
	c := service.C{
		BY: b,
	}
	return c
}