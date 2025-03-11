package product

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	useProduct "lab1/go-rest-api/internal/usecase/product"
)

func GetProductList(c *gin.Context) {

	productList, err := useProduct.GetProductList()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, productList)
}

type AddProductRequest struct {
	CustomerId uuid.UUID `json:"customer_id" binding:"required"`
	ProductId  uuid.UUID `json:"product_id" binding:"required"`
}

func AddToCart(c *gin.Context) {
	var req AddProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	cartId, err := useProduct.AddToCart(
		req.CustomerId,
		req.ProductId,
	)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, cartId)
}
