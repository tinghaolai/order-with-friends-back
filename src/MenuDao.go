package src

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func LoginCheck() gin.HandlerFunc {
    return func (c *gin.Context) {
        if _,status:=c.GetQuery("token");!status {
            c.String(http.StatusUnauthorized, "no token")
            c.Abort()
        } else {
            c.Next()
        }
    }
}

func GetMenu(c *gin.Context) {
    menu := Menu{}
    DBHelper.Find(&menu, c.Param("menu_id"))
    c.JSON(200, menu)
}

func GetMenus(c *gin.Context) {
    query:=MenuQuery{}
    err:=c.BindQuery(&query)
    if err!=nil {
        c.String(400, "param error %s", err.Error())
    } else {
        c.JSON(200, query)
    }
}

func StoreMenu(c *gin.Context) {
    menu:=MenuRequest{}
    err:=c.BindJSON(&menu)
    if err!=nil {
        c.String(400, "param error %s", err.Error())
    } else {
        c.JSON(200, menu)
    }

    c.String(200, "store response")
}
