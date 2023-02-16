package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func primeHandler(l *log.Logger) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			nums    []interface{}
			numsInt []int
		)

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

// toInt transform slice of interfaces to int slices and return error if one of entries is not int, err contain index of NaN
func toInt(nums []interface{}) (res []int, err error) {
	res = make([]int, len(nums))
	for i, num := range nums {
		switch v := num.(type) {
		case int:
			res[i] = v
		case float64:
			// as of all numbers in json unmarshalled as float, additional check required
			// if number eq itself after casting to int and then to float, we can tell what it is int
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
