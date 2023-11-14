package routers

import (
	"go-backend/handler"

	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	//router.Use(cors.Default)

	router.POST("/product", handler.InsertProduct)
	router.POST("/user", handler.InsertUser)
	router.GET("product/:id", handler.GetProduct)
	return router

}
