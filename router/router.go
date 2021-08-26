package router

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

func Init() {
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404, page not exists!",
		})
	})
	r.Run(":8000")
}
