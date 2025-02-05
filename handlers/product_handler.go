package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "t2/models"
)

var products = make(map[int]models.Product)
var lastID int

// GetProducts devuelve el mapa de productos
func GetProducts() map[int]models.Product {
    return products
}

func CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    lastID++
    product.ID = lastID
    products[product.ID] = product

    c.JSON(http.StatusCreated, product)
}

func GetProduct(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    product, exists := products[id]
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    var updatedProduct models.Product
    if err := c.ShouldBindJSON(&updatedProduct); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if _, exists := products[id]; !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    updatedProduct.ID = id
    products[id] = updatedProduct

    c.JSON(http.StatusOK, updatedProduct)
}

func DeleteProduct(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    if _, exists := products[id]; !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    delete(products, id)
    c.Status(http.StatusNoContent)
}