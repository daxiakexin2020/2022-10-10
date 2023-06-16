package handler

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewJsonHandler, NewBase, NewEnDeHandler, NewSymmetryEnDeHandler, NewComprehensiveHandler)
