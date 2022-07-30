package main

import (
	// "fmt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestModel struct{
	Id int `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		// c.String(http.StatusOK, "Hello world !!") // Returning a string

		// returning a JSON
		c.JSON(http.StatusOK, gin.H{
			"responseData" : "Hello Neeraj",
		})
	})

	r.GET("/:name", func(c *gin.Context) {
		
		var name = c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"responseData" : name,
		})
	})

	r.POST("/add", func(c *gin.Context){
		var data TestModel
		if err := c.ShouldBind(&data); err != nil{
			fmt.Printf("err: %s \n", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : fmt.Sprintf("%v", err),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data":data,
			})
		}
	})

	r.Run("localhost:8000")
}