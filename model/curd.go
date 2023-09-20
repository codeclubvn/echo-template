package model

type CRUD interface {
	Get(id string) (any, error)
	GetAll() (any, any, error)
	Create(data map[string]any) (any, error)
	Update(id string, data map[string]any) (any, error)
	Delete(id string) (any, error)
	Name() string
}
