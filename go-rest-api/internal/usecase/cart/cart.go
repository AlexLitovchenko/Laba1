package cart

import (
	"fmt"
	"lab1/go-rest-api/internal/models"

	"github.com/google/uuid"
)

func GetCart(customerId string) ([]models.OrderCartItem, error) {
	cartRepo := models.OrderCartRepo{}
	itemRepo := models.OrderCartItemRepo{}

	cart, err := cartRepo.GetCart(customerId)
	if err != nil || cart == nil {
		return nil, fmt.Errorf("no cart")
	}

	itemList, err := itemRepo.List(cart.ID)
	if err != nil {
		return nil, err
	}

	return itemList, nil
}

func CartAddProduct(itemId, cartId uuid.UUID) ([]models.OrderCartItem, error) {
	itemRepo := models.OrderCartItemRepo{}

	item, err := itemRepo.Get(itemId)
	if err != nil {
		return nil, err
	}

	item.Quantity++
	if err := itemRepo.Update(*item); err != nil {
		return nil, err
	}

	itemList, err := itemRepo.List(cartId)
	if err != nil {
		return nil, err
	}

	return itemList, nil
}

func CartDelProduct(itemId, cartId uuid.UUID) ([]models.OrderCartItem, error) {
	itemRepo := models.OrderCartItemRepo{}

	item, err := itemRepo.Get(itemId)
	if err != nil {
		return nil, err
	}

	item.Quantity--
	if item.Quantity > 0 {
		if err := itemRepo.Update(*item); err != nil {
			return nil, err
		}
	} else {
		if err := itemRepo.Delete(item.ID); err != nil {
			return nil, err
		}
	}

	itemList, err := itemRepo.List(cartId)
	if err != nil {
		return nil, err
	}
	return itemList, nil
}
