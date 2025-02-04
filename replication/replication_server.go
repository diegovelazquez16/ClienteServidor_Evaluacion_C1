package replication

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "sync"
    "t2/models"
)

var replicatedProducts = make(map[int]models.Product)
var mu sync.Mutex

func ReplicateProduct(product models.Product) {
    mu.Lock()
    defer mu.Unlock()
    replicatedProducts[product.ID] = product
}

func GetReplicatedProducts(c *gin.Context) {
    mu.Lock()
    defer mu.Unlock()

    products := make([]models.Product, 0, len(replicatedProducts))
    for _, product := range replicatedProducts {
        products = append(products, product)
    }

    c.JSON(http.StatusOK, products)
}