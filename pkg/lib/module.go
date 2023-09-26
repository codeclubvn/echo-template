package lib

import "go.uber.org/fx"

var Module = fx.Options(fx.Provide(NewZapLogger, NewServer, NewServerGroupV1), fx.Invoke(NewSwagger))
