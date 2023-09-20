package controller

import "template/model"

const DefaultCapacity = 5

type Controller struct {
	MapHandlers map[string]model.CRUD
}

func NewHandler(capacity ...int) *Controller {
	cap := DefaultCapacity
	if len(capacity) > 0 {
		cap = capacity[0]
	}
	handlers := make(map[string]model.CRUD, cap)
	return &Controller{
		MapHandlers: handlers,
	}
}
