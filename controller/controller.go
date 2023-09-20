package controller

import (
	"sync"

	"template/model"
)

type AppController struct {
	handlers *Controller
	sync.Mutex
}

func NewAppController() *AppController {
	ctrl := AppController{}
	ctrl.InitMaphandlers()

	return &ctrl
}

func (c *AppController) InitMaphandlers() {
	resources := c.NewResourceCRUD()

	h := NewHandler(len(resources))
	for _, i := range resources {
		// srcName := SingularToPlural(i.Name()) must convert resource name to singular
		h.MapHandlers[i.Name()] = i
	}

	c.handlers = h
}

func (c *AppController) NewResourceCRUD() []model.CRUD {
	return []model.CRUD{
		NewExample(c),
	}
}

func (c *AppController) GetMapHandlers() map[string]model.CRUD {
	return c.handlers.MapHandlers
}
