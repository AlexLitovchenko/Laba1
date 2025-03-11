package product

import (
	"lab1/go-rest-api/internal/models"

	"github.com/google/uuid"
)

func GetProductList() ([]models.Product, error) {
	productRepo := models.ProductRepo{}

	productList, err := productRepo.GetList()
	if err != nil {
		return nil, err
	}

	return productList, nil
}

func AddToCart(customerId, productId uuid.UUID) (uuid.UUID, error) {

	cartRepo := models.OrderCartRepo{}
	itemRepo := models.OrderCartItemRepo{}

	cart, err := cartRepo.GetCart(customerId.String())
	if err != nil {
		return uuid.UUID{}, err
	}

	var cartId uuid.UUID
	// карта есть, добавляем товар
	if cart == nil {
		cartNewId, err := cartRepo.CreateCart(
			models.OrderCart{
				ID:         uuid.New(),
				CustomerID: customerId,
			})
		if err != nil {
			return uuid.UUID{}, nil
		}
		cartId = cartNewId
	} else {
		cartId = cart.ID
	}

	// добавляем товар
	if err := itemRepo.Update(models.OrderCartItem{
		ID:          uuid.New(),
		OrderCartID: cartId,
		ProductID:   productId,
		Quantity:    1,
	}); err != nil {
		return uuid.UUID{}, err
	}

	return cartId, nil
}
