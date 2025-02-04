package main

import (
    "log"
    "time"
    "github.com/gin-gonic/gin"
    "t2/handlers"
    "t2/replication"
)

func main() {
    // Crear un router con Gin
    r := gin.Default()

    // CRUD endpoints
    r.POST("/products", handlers.CreateProduct)
    r.GET("/products/:id", handlers.GetProduct)
    r.PUT("/products/:id", handlers.UpdateProduct)
    r.DELETE("/products/:id", handlers.DeleteProduct)

    // Replication endpoint
    r.GET("/replicated-products", replication.GetReplicatedProducts)

    // Iniciar el servidor principal
    go func() {
        log.Println("Main server started at :8080")
        if err := r.Run(":8080"); err != nil {
            log.Fatalf("Failed to start main server: %v", err)
        }
    }()

    // Iniciar el servidor de replicación
    go func() {
        replicationRouter := gin.Default()
        replicationRouter.GET("/replicated-products", replication.GetReplicatedProducts)

        log.Println("Replication server started at :8081")
        if err := replicationRouter.Run(":8081"); err != nil {
            log.Fatalf("Failed to start replication server: %v", err)
        }
    }()

    // Simular short polling para replicación
    go func() {
        for {
            time.Sleep(5 * time.Second) // Poll cada 5 segundos
            for _, product := range handlers.Products {
                replication.ReplicateProduct(product)
            }
        }
    }()

    // Mantener la función main en ejecución
    select {}
}