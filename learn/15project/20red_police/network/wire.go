package network

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(Gresources)
