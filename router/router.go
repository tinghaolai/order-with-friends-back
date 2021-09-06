package router

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	. "owf/src"
)

func Init() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
        v1.GET("/menu/:menu_id", GetMenu)
        v1.GET("/menus", GetMenus)
        v1.Use(LoginCheck())
        {
            v1.POST("/menu", StoreMenu)
        }
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404, page not exists!",
		})
	})
	r.Run(":8000")
}
