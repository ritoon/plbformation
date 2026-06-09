package postgres

import (
	"plbformation/interface/4/db"
	"plbformation/interface/4/model"

	"gorm.io/gorm"
)

var (
	_ db.UserStore = (*PostgresUser)(nil)
)

type PostgresUser struct {
	db *gorm.DB
}

func (p *PostgresUser) Create(user model.User) error {
	return p.db.Create(&user).Error
}

func (p *PostgresUser) Delete(id int) error {
	return p.db.Delete(&model.User{}, id).Error
}

// Implementation of UserStore interface
func (p *PostgresUser) Save(user *model.User) error {
	return p.db.Save(user).Error
}

func (p *PostgresUser) GetByID(id int) (*model.User, error) {
	var user model.User
	if err := p.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (p *PostgresUser) GetAll() ([]model.User, error) {
	var users []model.User
	if err := p.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
