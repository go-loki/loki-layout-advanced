package server

import (
	"github.com/google/wire"
)

// ProviderHttpServerSet is server providers.
var ProviderHttpServerSet = wire.NewSet(NewHttpServer)
