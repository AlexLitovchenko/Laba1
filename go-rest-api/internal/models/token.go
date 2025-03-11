package models

import (
	"lab1/go-rest-api/internal/db"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AccessToken - токен доступа (1:1 с Customer)
type AccessToken struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CustomerID uuid.UUID `gorm:"type:uuid;uniqueIndex;not null"`
	Token      string    `gorm:"type:varchar(255);not null"`
	IsActive   bool      `gorm:"default:true"`
	CreatedAt  time.Time `gorm:"not null"`
	UpdatedAt  time.Time `gorm:"not null"`
}

type AccessTokenRepo struct{}

func (self *AccessTokenRepo) TableName() string {
	return "access_tokens"
}

func (self *AccessTokenRepo) GetDb() *gorm.DB {
	return db.GetDb().Table(self.TableName())
}

func (self *AccessTokenRepo) Add(token AccessToken) error {
	req := self.GetDb().Save(&token)
	if req.Error != nil {
		return req.Error
	}
	return nil
}

func (self *AccessTokenRepo) Find(token string) (bool, error) {
	var access AccessToken
	req := self.GetDb().First(&access).Where("token = ?", token)
	if req.Error != nil {
		if req.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, req.Error
	}
	return true, nil
}
