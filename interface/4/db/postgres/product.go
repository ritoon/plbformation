package postgres

import (
	"plbformation/interface/4/db"
	"plbformation/interface/4/model"

	"gorm.io/gorm"
)

var (
	_ db.ProductStore = (*PostgresProduct)(nil)
)

type PostgresProduct struct {
	db *gorm.DB
}

func (p *PostgresProduct) Create(product *model.Product) error {
	if product == nil {
		return db.NewErrorDb("product is nil", nil, "Create Product", true)
	}

	return p.db.Create(&product).Error
}

func (p *PostgresProduct) Delete(id int) error {
	return p.db.Delete(&model.Product{}, id).Error
}

// Implementation of ProductStore interface
func (p *PostgresProduct) Save(product *model.Product) error {
	return p.db.Save(product).Error
}

func (p *PostgresProduct) GetByID(id int) (*model.Product, error) {
	var product model.Product
	if err := p.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *PostgresProduct) GetAll() ([]model.Product, error) {
	var products []model.Product
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
