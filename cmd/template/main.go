package main

import (
	apiV1 "template/api"
	"template/controller"
	"template/internal"
	"template/system"
)

func main() {
	ctrl := controller.NewAppController()
	LoadApplicationApi(ctrl)
}

func LoadApplicationApi(ctrl *controller.AppController) error {
	server, err := system.NewHTTPInstance()
	if err != nil {
		return err
	}

	appApiV1 := apiV1.NewApiService()

	internal.ApplyHandler(ctrl, appApiV1)

	go appApiV1.Run(server)
	return nil
}
