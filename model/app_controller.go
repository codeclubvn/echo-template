package model

type AppController interface {
	GetMapHandlers() map[string]CRUD
}
