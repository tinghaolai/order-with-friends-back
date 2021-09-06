package src

type Menu struct {
    MenuID int `json:"id"`
    MenuTitle string `json:"title" binding:"required"`
    MenuItemList []MenuItem `json:"items" binding:"lt=100,multiItemSize,dive"`
    MenuItemSize int `json:"itemSize"`
}

type MenuItem struct {
    ItemName string `json:"name" binding:"required"`
    Price int `json:"price" binding:"required,gt=0"`
}

type MenuQuery struct {
    Page int `json:"page" form:"page" binding:"required"`
    PageSize int `json:"size" form:"size" binding:"required"`
}
