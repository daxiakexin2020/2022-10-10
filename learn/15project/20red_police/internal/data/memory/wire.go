package memory

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewUser, NewPlayer)
