package service

import (
	"github.com/google/wire"
)

// ProviderServiceSet is service providers.
var ProviderServiceSet = wire.NewSet(
	NewIndexService,
)
