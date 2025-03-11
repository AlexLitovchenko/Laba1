package handler

import (
	"lab1/go-rest-api/internal/handler/api/auth"
	"lab1/go-rest-api/internal/handler/api/cart"
	"lab1/go-rest-api/internal/handler/api/product"
	"lab1/go-rest-api/internal/handler/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes регистрирует маршруты и возвращает роутер
func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Регистрация обработчиков
	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)

	authRoutes := r.Group("/").Use(middleware.AuthMiddleware())
	authRoutes.GET("/cart", cart.GetCart)
	authRoutes.GET("/product/list", product.GetProductList)
	authRoutes.POST("/product/add", product.AddToCart)
	authRoutes.POST("/cart/product-add", cart.Add)
	authRoutes.POST("/cart/product-del", cart.Del)

	return r
}
