package lib

import "go.uber.org/fx"

var Module = fx.Options(fx.Provide(NewZapLogger, NewRequestValidator, NewServer, NewServerGroupV1), fx.Invoke(NewSwagger))
