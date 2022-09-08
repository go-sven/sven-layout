package index

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
		NewDemoHandler,
	)