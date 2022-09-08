package router

import "github.com/google/wire"

var ProviderSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)) )
