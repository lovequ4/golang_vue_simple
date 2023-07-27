package services

import (
	"app/database"
	"app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func CreateOrder(c *gin.Context) {
	var requestBody struct {
		MenuItem []struct {
			Menu     models.Menu `json:"menu"`
			Quantity int         `json:"quantity"`
		} `json:"menuitem"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var totalOrderPrice int

	var menuItems []models.MenuItem
	for _, menuItemData := range requestBody.MenuItem {
		subtotal := menuItemData.Menu.Price * menuItemData.Quantity
		totalOrderPrice += subtotal

		menuItem := models.MenuItem{
			Menu:     menuItemData.Menu,
			Quantity: menuItemData.Quantity,
			Subtotal: subtotal,
		}
		menuItems = append(menuItems, menuItem)
	}

	order := models.Order{
		MenuItem:   menuItems,
		TotalPrice: totalOrderPrice,
		CreateTime: time.Now(),
	}

	// Save the order to the database
	result := database.DB.Create(&order)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func GetOrder(c *gin.Context) {
	orderId := c.Param("orderId")

	var order models.Order

	err := database.DB.Preload("MenuItem.Menu").First(&order, "id = ?", orderId).Error
	if err != nil {
		// Handle the error if the order is not found or other database errors occur
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	for i := range order.MenuItem {
		order.MenuItem[i].Menu.MenuImg = "http://" + c.Request.Host + "/" + order.MenuItem[i].Menu.MenuImg
	}

	// Return the order data as a response
	c.JSON(http.StatusOK, order)
}

func CheckOrder(c *gin.Context) {
	orderID := c.Param("orderId")

	paymentURL := "http://" + c.Request.Host + "/" + orderID

	qrCode, err := qrcode.New(paymentURL, qrcode.Medium)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code"})
		return
	}

	imageBytes, err := qrCode.PNG(512)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code image"})
		return
	}

	c.Data(http.StatusOK, "image/png", imageBytes)
}
