package db

import "plbformation/interface/4/model"

type DBStore struct {
	User    UserStore
	Product ProductStore
}

type UserStore interface {
	Save(user *model.User) error
	GetByID(id int) (*model.User, error)
	GetAll() ([]model.User, error)
	Create(user *model.User) error
	Delete(id int) error
}

type ProductStore interface {
	Save(product *model.Product) error
	GetByID(id int) (*model.Product, error)
	GetAll() ([]model.Product, error)
	Create(product *model.Product) error
	Delete(id int) error
}
