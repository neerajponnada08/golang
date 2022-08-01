package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestModel struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding "required"`
}

func main() {
	r := gin.Default()
	v1 := r.Group("v1")
	{

		user := v1.Group("user")
		{
			user.GET("", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"data": []TestModel{
						TestModel{
							Id:   1,
							Name: "Neeraj",
						},
						TestModel{
							Id:   2,
							Name: "Sindhu",
						},
					},
				})
			})

			user.GET(":Id", func(c *gin.Context) {
				var id = c.Param("Id")
				fmt.Println("Id is received: ", id)
				c.JSON(http.StatusOK, gin.H{
					"data": TestModel{
						Id:   1,
						Name: "Neeraj",
					},
				})
			})
		}
		product := v1.Group("product")
		{
			product.GET(":id", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"data": TestModel{
						Id:   1,
						Name: "Product1",
					},
				})
			})
		}
	} //v1 ending

	v2 := r.Group("v2")
	{

		user := v2.Group("user")
		{
			user.GET("", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"data": []TestModel{
						TestModel{
							Id:   1,
							Name: "v2Neeraj",
						},
						TestModel{
							Id:   2,
							Name: "v2Sindhu",
						},
					},
				})
			})

			user.GET(":Id", func(c *gin.Context) {
				var id = c.Param("Id")
				fmt.Println("Id is received: ", id)
				c.JSON(http.StatusOK, gin.H{
					"data": TestModel{
						Id:   1,
						Name: "v2Neeraj",
					},
				})
			})
		}
		product := v2.Group("v2product")
		{
			product.GET("", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"data": []TestModel{
						TestModel{
							Id:   1,
							Name: "v2Procutc1",
						},
						TestModel{
							Id:   2,
							Name: "v2Product2",
						},
					},
				})
			})
		}
	}

	r.Run("localhost:8000")
}
