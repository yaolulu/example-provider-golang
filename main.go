package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/products/:id", GetProduct)
	router.Run("localhost:8080")
}
