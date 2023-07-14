package server

import (
	"github.com/google/wire"
)

// ProviderHandlerSet is handler providers.
var ProviderHandlerSet = wire.NewSet(
	NewIndexHandler,
)
