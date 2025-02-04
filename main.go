package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "t2/handlers"
)

func main() {
    r := gin.Default()

    r.POST("/products", handlers.CreateProduct)
    r.GET("/products/:id", handlers.GetProduct)
    r.PUT("/products/:id", handlers.UpdateProduct)
    r.DELETE("/products/:id", handlers.DeleteProduct)


    go func() {
        log.Println("Main server started at :8080")
        if err := r.Run(":8080"); err != nil {
            log.Fatalf("Failed to start main server: %v", err)
        }
    }()

    go func() {
        replicationRouter := gin.Default()

        log.Println("Replication server started at :8081")
        if err := replicationRouter.Run(":8081"); err != nil {
            log.Fatalf("Failed to start replication server: %v", err)
        }
    }()



    select {}
}