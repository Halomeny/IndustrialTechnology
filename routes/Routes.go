package routes

import (
	"indpr/Secrets"
	"indpr/productsActions"
	"indpr/userActions"

	// Gin - для развертка API

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Настройка путей
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/login", userActions.Login)
	r.POST("/register", userActions.Register)
	r.GET("/getallproducts", productsActions.GetAllProducts)

	protected := r.Group("/")
	protected.Use(Secrets.AuthMiddleware())
	{
		protected.POST("/addproduct", productsActions.AddProduct)
		protected.POST("/delorrestoreproduct", productsActions.DelOrRestoreProduct)
	}

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
