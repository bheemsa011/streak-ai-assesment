package controllers

import (
	"net/http"
	"streak-ai-assesment/model"

	"github.com/gin-gonic/gin"
)

func FindPairs(c *gin.Context) {

	var request model.Request

	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request input"})
		return
	}

	solutions := [][]int{}
	seen := make(map[int][]int)

	for index, num := range request.Numbers {
		complement := request.Target - num
		if indices, found := seen[complement]; found {
			for _, i := range indices {
				solutions = append(solutions, []int{i, index})
			}
		}
		seen[num] = append(seen[num], index)
	}

	c.JSON(http.StatusOK, model.Response{Solutions: solutions})
}
