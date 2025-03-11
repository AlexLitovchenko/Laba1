package customer

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"lab1/go-rest-api/internal/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(
	name, lastName, email, password string,
) error {

	customerRepo := models.CustomerRepo{}

	err := customerRepo.Add(models.Customer{
		ID:       uuid.New(),
		Name:     name,
		LastName: lastName,
		Email:    email,
		Password: password,
	})

	if err != nil {
		return err
	}
	return nil
}

func Login(email, password string) (models.AccessToken, error) {
	customerRepo := models.CustomerRepo{}
	tokenRepo := models.AccessTokenRepo{}

	customer, err := customerRepo.Get(email)
	if err != nil {
		return models.AccessToken{}, err
	}

	if customer.Password != password {
		return models.AccessToken{}, fmt.Errorf("wrong password")
	}

	token, err := generateToken(email)
	if err != nil {
		return models.AccessToken{}, err
	}

	access := models.AccessToken{
		ID:         uuid.New(),
		CustomerID: customer.ID,
		Token:      token,
		IsActive:   true,
	}

	err = tokenRepo.Add(access)
	if err != nil {
		return models.AccessToken{}, err
	}

	return access, nil
}

func generateToken(email string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
