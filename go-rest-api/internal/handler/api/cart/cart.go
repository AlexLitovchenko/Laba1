package cart

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	useCart "lab1/go-rest-api/internal/usecase/cart"
)

func GetCart(c *gin.Context) {
	customerId := c.Query("customer_id")

	cartItems, err := useCart.GetCart(customerId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, cartItems)
}

type CartUpdateRequest struct {
	CartItemId uuid.UUID `json:"cart_item_id" binding:"required"`
	CartId     uuid.UUID `json:"cart_id" binding:"required"`
}

func Add(c *gin.Context) {
	var req CartUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	cart, err := useCart.CartAddProduct(req.CartItemId, req.CartId)
	if err != nil || cart == nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, cart)
}

func Del(c *gin.Context) {
	var req CartUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	cart, err := useCart.CartDelProduct(req.CartItemId, req.CartId)
	if err != nil || cart == nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, cart)
}
