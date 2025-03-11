package models

import (
	"lab1/go-rest-api/internal/db"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Product - товар (1:1 с OrderCartItem)
type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title       string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	Price       int       `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductRepo struct{}

func (*ProductRepo) TableName() string {
	return "products"
}

func (self *ProductRepo) GetDb() *gorm.DB {
	return db.GetDb().Table(self.TableName())
}

func (self *ProductRepo) GetList() ([]Product, error) {
	var productList []Product
	find := self.GetDb().Find(&productList)
	if find.Error != nil {
		return nil, find.Error
	}

	return productList, nil
}
