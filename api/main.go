package main

import (
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"

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

	// route
	router.POST("/process-single", processSingle)
	router.POST("/process-concurrent", processConcurrent)

	fmt.Printf("Server running on port: %d", port)
	router.Run(fmt.Sprintf(":%d", port))
}

// ---------process single sequentially---
func processSingle(c *gin.Context) {

	var req SortRequest

	// c.ShouldBindJSON to bind the json data in user struct data formate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// calculating time duration
	startTime := time.Now()

	// performing sorting
	sortedArrays := make([][]int, len(req.ToSort))
	for i, arr := range req.ToSort {

		sortedArr := make([]int, len(arr))
		copy(sortedArr, arr)
		sort.Ints(sortedArr)
		sortedArrays[i] = sortedArr
	}

	duration := time.Since(startTime)
	// bind result into struct
	res := SortResponse{
		SortedArrays: sortedArrays,
		TimeNS:       duration.Nanoseconds(),
	}

	// response struct send in response in JSON formate
	c.JSON(http.StatusOK, res)

}

// --------process concurrent
func processConcurrent(c *gin.Context) {
	var req SortRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	startTime := time.Now()

	// sync using fo concurrent processing
	var wg sync.WaitGroup
	var mu sync.Mutex

	sortedArrays := make([][]int, len(req.ToSort))

	// sorting
	for i, arr := range req.ToSort {
		wg.Add(1)
		go func(i int, arr []int) {
			defer wg.Done()
			sortedArr := make([]int, len(arr))
			copy(sortedArr, arr)
			sort.Ints(sortedArr)

			mu.Lock()
			sortedArrays[i] = sortedArr
			mu.Unlock()
		}(i, arr)
	}
	wg.Wait()

	duration := time.Since(startTime)
	// data binding
	res := SortResponse{
		SortedArrays: sortedArrays,
		TimeNS:       duration.Nanoseconds(),
	}

	// send response
	c.JSON(http.StatusOK, res)

}
