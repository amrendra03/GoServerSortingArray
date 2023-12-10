package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting...")

	router := gin.Default()

	port := 8000

	fmt.Printf("Server running on port: %d", port)

	router.Run(fmt.Sprintf(":%d", port))
}
