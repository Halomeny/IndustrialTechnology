package productsActions

import (
	"fmt"
	"indpr/dataBase"
	models "indpr/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CredentialsForAddProduct struct {
	Name        string  `json:"Название"`
	Price       float32 `json:"цена"`
	Description string  `json:"Описание"`
}
type CredentialsForDelProduct struct {
	Name         string `json:"Название"`
	DeletionMark bool   `json:"Пометка удаления"`
}

// @Summary      Add new product
// @Description  Adding a product with a name and price
// @Security     BearerAuth
// @Accept      json
// @Produce      json
// @Param      body body CredentialsForAddProduct true "Product"
// @Success    200 "Новый товар успешно добавлен"
// @Router       /addproduct [post]
// @Tags ProductsActions
func AddProduct(c *gin.Context) {
	var creds CredentialsForAddProduct
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Недопустимый запрос"})
		return
	}

	product := models.Product{
		Name:         creds.Name,
		Price:        float32(creds.Price),
		Description:  creds.Description,
		DeletionMark: false,
	}

	if err := checkExistenceProduct(product.Name); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := dataBase.Db.Create(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// @Summary      Delete or restore product
// @Description  Allows you to remove or mark the deletion
// @Security     BearerAuth
// @Accept      json
// @Produce      json
// @Param      body body CredentialsForDelProduct true "Product"
// @Success    200 "Пометка удаления успешно изменена"
// @Router       /delorrestoreproduct [post]
// @Tags ProductsActions
func DelOrRestoreProduct(c *gin.Context) {
	var creds CredentialsForDelProduct
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Недопустимый запрос"})
		return
	}

	product := models.Product{
		Name:         creds.Name,
		DeletionMark: creds.DeletionMark,
	}

	if err := checkExistenceProduct(product.Name); err == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := dataBase.Db.Model(&models.Product{}).Where("name = ?", product.Name).Updates(product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary      Get products
// @Description  Returns a list of all products
// @Router       /getallproducts [get]
// @Tags ProductsActions
func GetAllProducts(c *gin.Context) {
	var products []models.Product
	dataBase.Db.Find(&products)
	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func checkExistenceProduct(productName string) error {
	var product models.Product

	if err := dataBase.Db.Where("name = ?", productName).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}

	log.Printf("ERROR: продукт с таким именем уже существет %s", productName)
	return fmt.Errorf("продукт с таким именем уже существет")

}
