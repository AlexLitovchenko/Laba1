package models

import (
	"lab1/go-rest-api/internal/db"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// OrderCart - корзина заказов (1:1 к Customer)
type OrderCart struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CustomerID uuid.UUID `gorm:"type:uuid;not null;index"`
	CreatedAt  time.Time `gorm:"not null"`
	UpdatedAt  time.Time `gorm:"not null"`
}
type OrderCartRepo struct{}

func (*OrderCartRepo) TableName() string {
	return "order_carts"
}

func (self *OrderCartRepo) GetDb() *gorm.DB {
	return db.GetDb().Table(self.TableName())
}

func (self *OrderCartRepo) GetCart(customerId string) (*OrderCart, error) {
	var cart OrderCart
	find := self.GetDb().First(&cart, "customer_id = ?", customerId)
	if find.Error != nil {
		if find.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, find.Error
	}

	return &cart, nil
}

func (self *OrderCartRepo) CreateCart(cart OrderCart) (uuid.UUID, error) {

	find := self.GetDb().Create(&cart)
	if find.Error != nil {
		return uuid.UUID{}, find.Error
	}

	return cart.ID, nil
}

func (self *OrderCartRepo) DeleteCart(cartId uuid.UUID) error {

	find := self.GetDb().Delete(&OrderCart{}, "id = ?", cartId)
	if find.Error != nil {
		return find.Error
	}

	return nil
}
