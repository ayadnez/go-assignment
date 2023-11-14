package handler

import (
	"go-backend/models"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InsertProduct(c *gin.Context) {
	db, err := models.InitDB()
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	//defer db.Close()
	var product models.Product

	// Bind JSON request to the Product struct
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	newProduct := &models.DbProduct{

		//ProductID:               product.ProductID,
		ProductName:             product.ProductName,
		ProductDescription:      product.ProductDescription,
		ProductImage:            models.ToString(product.ProductImage),
		ProductPrice:            product.ProductPrice,
		CompressedProductImages: models.ToString(product.CompressedProductImages),
		CreatedAt:               time.Now(),
		UpdatedAt:               time.Now(),
	}

	db.Create(&newProduct)

	go func() {
		var wg sync.WaitGroup
		wg.Add(1)
		Downloadimg(product.ProductImage, newProduct.ProductID)
		wg.Wait()

	}()

	// Respond with the created product
	c.JSON(200, gin.H{
		"status": "success",
		"result": "inserted the product",
	})
}

func InsertUser(c *gin.Context) {
	db, err := models.InitDB()
	if err != nil {
		panic("failed to connect to database" + err.Error())
	}
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": "invalid request payload"})
		return

	}
	newuser := &models.User{
		Id:         user.Id,
		Name:       user.Name,
		Mobile:     user.Mobile,
		Latitude:   user.Latitude,
		Longitude:  user.Longitude,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	db.Create(&newuser)
	c.JSON(200, gin.H{
		"status:": "scucess",
		"result:": "inserted the user",
	})
}

func GetProduct(c *gin.Context) {
	db, err := models.InitDB()
	if err != nil {
		panic("faled to connect to database" + err.Error())
	}
	var id = c.Param("id")
	var product models.DbProduct
	result := db.First(&product, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{

			"status":  "error",
			"message": "prdouct not found",
		})
		return
	}

	c.JSON(200, gin.H{

		"status": "success",
		"result": product,
	})

}
