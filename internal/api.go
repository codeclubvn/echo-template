package internal

import (
	"template/api"
	"template/model"
)

func ApplyHandler(svc model.AppController, c *api.Api) {
	// router here
}

func GetHandlers(svc model.AppController, resource string) (model.CRUD, error) {
	mapHandler := svc.GetMapHandlers()
	src, ok := mapHandler[resource]
	if !ok {
		// handler error
	}
	return src, nil
}
