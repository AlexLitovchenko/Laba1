package models

import (
	"lab1/go-rest-api/internal/db"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// OrderCartItem - товары в корзине (1:1 с Product)
type OrderCartItem struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OrderCartID uuid.UUID `gorm:"type:uuid;not null;index"`
	ProductID   uuid.UUID `gorm:"type:uuid;not null"`
	Quantity    int       `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}

type OrderCartItemRepo struct{}

func (self *OrderCartItemRepo) TableName() string {
	return "order_cart_items"
}

func (self *OrderCartItemRepo) GetDb() *gorm.DB {
	return db.GetDb().Table(self.TableName())
}

func (self *OrderCartItemRepo) Update(item OrderCartItem) error {
	req := self.GetDb().Save(&item)
	if req.Error != nil {
		return req.Error
	}
	return nil
}

func (self *OrderCartItemRepo) Delete(itemId uuid.UUID) error {
	req := self.GetDb().Delete(&OrderCartItem{}, "id = ?", itemId)
	if req.Error != nil {
		return req.Error
	}
	return nil
}

func (self *OrderCartItemRepo) Get(itemId uuid.UUID) (*OrderCartItem, error) {
	var item OrderCartItem

	req := self.GetDb().Find(&item).Where("id = ?", itemId)
	if req.Error != nil {
		return nil, req.Error
	}

	return &item, nil
}

func (self *OrderCartItemRepo) List(cartId uuid.UUID) ([]OrderCartItem, error) {

	var list []OrderCartItem

	req := self.GetDb().Find(&list).Where("order_cart_id = ?", cartId)
	if req.Error != nil && req.Error != gorm.ErrRecordNotFound {
		return nil, req.Error
	}

	return list, nil
}
