package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type SortRequest struct {
	ToSort [][]int `json:"to_sort"`
}

type SortResponse struct {
	SortedArrays [][]int `json:"sorted_arrays"`
	TimeNS       int64   `json:"time_ns"`
}

func main() {
	fmt.Println("Starting...")

	router := gin.Default()

	port := 8000

	fmt.Printf("Server running on port: %d", port)

	router.Run(fmt.Sprintf(":%d", port))
}
