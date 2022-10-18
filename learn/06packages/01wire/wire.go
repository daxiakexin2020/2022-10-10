package main

import (
	"01wire/service"
	"github.com/google/wire"
)

func Initialize(msg string) service.C {
	panic(wire.Build(service.C{}, service.B{}, service.A{}))
}
