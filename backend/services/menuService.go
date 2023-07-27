package services

import (
	"app/database"
	"app/models"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ServeMenuImage(c *gin.Context) {
	imagePath := c.Param("imagePath")

	// Get the absolute path to the image file
	absPath := filepath.Join("menusimg", imagePath)

	c.File(absPath)
}

func FindAllMenu(c *gin.Context) {
	var menus []models.Menu

	result := database.DB.Find(&menus)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	for i := range menus {
		menus[i].MenuImg = "http://" + c.Request.Host + "/" + menus[i].MenuImg
	}
	c.JSON(http.StatusOK, menus)
}

func CreateMenu(c *gin.Context) {
	menuname := c.Request.FormValue("menuname")
	description := c.Request.FormValue("description")

	//formdata價格字串轉數字
	priceStr := c.Request.FormValue("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid price",
		})
		return
	}

	//formdata圖片處理
	var menuImgPath string
	imageFile, header, err := c.Request.FormFile("menuimg")
	if err == nil {
		imageData, err := ioutil.ReadAll(imageFile)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to read image file",
			})
			return
		}

		fileName := header.Filename
		filePath := filepath.Join("menusimg", fileName)

		err = ioutil.WriteFile(filePath, imageData, 0644)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to save image file",
			})
			return
		}
		menuImgPath = filePath
	}

	newMenu := models.Menu{
		MenuName:    menuname,
		Description: description,
		Price:       price,
		MenuImg:     menuImgPath,
	}

	result := database.DB.Create(&newMenu)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create Menu",
		})
	}

	c.JSON(http.StatusOK, newMenu)
}
