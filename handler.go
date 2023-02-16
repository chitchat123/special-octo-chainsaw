package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func primeHandler(l *log.Logger) func(c *gin.Context) {
	return func(c *gin.Context) {
		var nums []interface{}
		var numsInt []int

		err := c.ShouldBindJSON(&nums)
		if err != nil {
			l.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("could not parse body: %w", err).Error()})
			c.Abort()
			return
		}
		// could be more optimal just bind to slice of int`s
		numsInt, err = toInt(nums)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("the given input is invalid. Element on index %s is not a number", err.Error())})
			c.Abort()
			return
		}

		var response = make([]bool, len(nums))
		for i, num := range numsInt {
			response[i] = cachedIsPrime(num)
		}

		c.JSON(http.StatusOK, response)
	}
}

func toInt(nums []interface{}) (res []int, err error) {
	res = make([]int, len(nums))
	for i, num := range nums {
		switch v := num.(type) {
		case int:
			res[i] = v
		case float64:
			if float64(int(v)) != v {
				return nil, fmt.Errorf("%d", i)
			}
			res[i] = int(v)
		default:
			return nil, fmt.Errorf("%d", i)
		}
	}
	return res, nil
}
