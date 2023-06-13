package local

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewJsonRepository, NewEnDeRepository, NewSymmetryEnDeRepository, NewComprehensiveRepository)
