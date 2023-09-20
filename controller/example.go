package controller

import "template/model"

const ResourceExample = "example"

type Example struct {
	r model.AppController
}

func NewExample(r model.AppController) *Example {
	return &Example{
		r: r,
	}
}

func (c *Example) Name() string {
	return ResourceExample
}

func (c *Example) Create(data map[string]any) (any, error) {
	return "created", nil
}

func (c *Example) Update(id string, data map[string]any) (any, error) {
	return "updated", nil
}

func (c *Example) Get(id string) (any, error) {
	return "successfully", nil
}

func (c *Example) Delete(id string) (any, error) {
	return "successfully", nil
}

func (c *Example) GetAll() (any, any, error) {
	return "successfully", nil, nil
}
