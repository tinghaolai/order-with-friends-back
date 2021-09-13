package src

type MenuRequest struct {
    MenuID int `json:"id"`
    MenuTitle string `json:"title" binding:"required"`
    MenuItemList []MenuItemRequest `json:"items" binding:"lt=100,multiItemSize,dive"`
    MenuItemSize int `json:"itemSize"`
}

type MenuItemRequest struct {
    ItemName string `json:"name" binding:"required"`
    Price int `json:"price" binding:"omitempty,gt=0"`
}

type MenuQuery struct {
    Page int `json:"page" form:"page" binding:"required"`
    PageSize int `json:"size" form:"size" binding:"required"`
}

type Menu struct {
    Id int `gorm:"primaryKey"`
    Title string
}
