package main

import (
    "log"
    "time"
    "github.com/gin-gonic/gin"
    "t2/handlers"
    "t2/replication"
)

func main() {
    r := gin.Default()

    r.POST("/products", handlers.CreateProduct)
    r.GET("/products/:id", handlers.GetProduct)
    r.PUT("/products/:id", handlers.UpdateProduct)
    r.DELETE("/products/:id", handlers.DeleteProduct)

    r.GET("/replicated-products", replication.GetReplicatedProducts)

    go func() {
        log.Println("Main server started at :8080")
        if err := r.Run(":8080"); err != nil {
            log.Fatalf("Failed to start main server: %v", err)
        }
    }()

    go func() {
        replicationRouter := gin.Default()
        replicationRouter.GET("/replicated-products", replication.GetReplicatedProducts)

        log.Println("Replication server started at :8081")
        if err := replicationRouter.Run(":8081"); err != nil {
            log.Fatalf("Failed to start replication server: %v", err)
        }
    }()

    go func() {
        for {
            time.Sleep(5 * time.Second) 
            products := handlers.GetProducts() 
            for _, product := range products {
                replication.ReplicateProduct(product)
            }
        }
    }()

    select {}
}