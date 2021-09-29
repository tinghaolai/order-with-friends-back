package router

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"net/http"
	. "owf/src"
	"log"
	"context"
	"time"
)

func Init() {
	r := gin.Default()
    corsConf := cors.DefaultConfig()
    corsConf.AllowAllOrigins = true
    corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
    corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin", "token", "Connection",
        "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"}

    r.Use(cors.New(corsConf))
	v1 := r.Group("/v1", EnableCookieSession())
	{
        v1.POST("/register", UserRegister)
        v1.POST("/login", UserLogin)
        v1.Use(LoginCheck())
        {
            v1.GET("/menu/:menu_id", GetMenu)
            v1.GET("/menus", GetMenus)
            v1.POST("/menu", StoreMenu)
        }
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404, page not exists!",
		})
	})

	server := &http.Server {
	    Addr: ":8000",
	    Handler: r,
	}

    go(func() {
    	err := server.ListenAndServe()
    	if err != nil {
    	    log.Fatal("Server error")
    	}
    })()

    SeverNotify()
    ctx,cancel:=context.WithTimeout(context.Background(), time.Second*5)
    defer cancel()
    err:=server.Shutdown(ctx)
    if err!=nil {
        log.Fatal("Server shut down")
    }
}
