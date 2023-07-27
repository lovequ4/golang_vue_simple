package main

import (
	"app/database"
	"app/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	database.DBconnect()
}

func main() {
	r := gin.Default()

	// cors
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	//menu
	r.GET("/menu", services.FindAllMenu)
	r.GET("/menusimg/:imagePath", services.ServeMenuImage)
	r.POST("/addmenu", services.CreateMenu)
	//order
	r.POST("/addorder", services.CreateOrder)
	r.GET("/order/:orderId", services.GetOrder)
	r.GET("/pay/:orderId", services.CheckOrder)
	r.Run()
}
