package controller

import "go.uber.org/fx"

var Module = fx.Options(fx.Provide(NewAuthController, NewUserController, NewPostController, NewFileController, NewImageController))
