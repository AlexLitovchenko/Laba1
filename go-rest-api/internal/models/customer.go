package models

import (
	"lab1/go-rest-api/internal/db"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Customer - клиент
type Customer struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time `gorm:"not null"`
	Name      string    `gorm:"type:varchar(50);not null"`
	LastName  string    `gorm:"type:varchar(50);not null"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string    `gorm:"type:varchar(50);not null"`
}

type CustomerRepo struct{}

func (*CustomerRepo) TableName() string {
	return "customers"
}

func (self *CustomerRepo) GetDb() *gorm.DB {
	return db.GetDb().Table(self.TableName())
}

func (self *CustomerRepo) Add(customer Customer) error {
	req := self.GetDb().Save(&customer)
	if req.Error != nil {
		return req.Error
	}
	return nil
}

func (self *CustomerRepo) Get(email string) (Customer, error) {
	var customer Customer
	req := self.GetDb().First(&customer).Where("email = ?", email)
	if req.Error != nil {
		return Customer{}, req.Error
	}
	return customer, nil
}
