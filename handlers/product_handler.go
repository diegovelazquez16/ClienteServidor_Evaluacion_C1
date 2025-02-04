package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "t2/models"
)

var products = make(map[int]models.Product)
var lastID int

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

