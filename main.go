package main

import (
	"github.com/gin-gonic/gin"
	productcontroller "github.com/raihannooriman/go-gin-gorm-restapi/controllers/productController"
	"github.com/raihannooriman/go-gin-gorm-restapi/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/products/:id", productcontroller.Show)
	r.POST("/api/products", productcontroller.Create)
	r.PUT("/api/products/:id", productcontroller.Update)
	r.DELETE("/api/products", productcontroller.Delete)

	r.Run()
}
