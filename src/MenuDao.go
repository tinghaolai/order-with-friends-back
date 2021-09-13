package src

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "gorm.io/gorm"
)

func LoginCheck() gin.HandlerFunc {
    return func (c *gin.Context) {
        if token :=c.GetHeader("token");token != "testToken" {
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
        var menuItems []MenuItem
          for i := 0; i < len(menu.MenuItemList); i++ {
              menuItems = append(menuItems, MenuItem {
                  Name: menu.MenuItemList[i].ItemName,
                  Price: menu.MenuItemList[i].Price,
                  Remark: menu.MenuItemList[i].Remark})
          }

        menuCreate := Menu {
            Title: menu.MenuTitle,
            Phone: menu.Phone,
            Remark: menu.Remark,
            MenuItems: menuItems}

        DBHelper.Transaction(func(tx *gorm.DB) error {
            return tx.Create(&menuCreate).Error
        })

        c.JSON(200, menu)
    }
}
