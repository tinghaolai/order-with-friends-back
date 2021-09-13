package src

type MenuRequest struct {
    MenuID int `json:"id"`
    MenuTitle string `json:"title" binding:"required"`
    MenuItemList []MenuItemRequest `json:"items" binding:"lt=100,multiItemSize,dive"`
    MenuItemSize int `json:"itemSize"`
    Phone string `json:"phone"`
    Remark string `json:"remark"`
}

type MenuItemRequest struct {
    ItemName string `json:"name" binding:"required"`
    Price int `json:"price" binding:"omitempty,gt=0"`
    Remark string `json:"remark"`
}

type MenuQuery struct {
    Page int `json:"page" form:"page" binding:"required"`
    PageSize int `json:"size" form:"size" binding:"required"`
}

type Menu struct {
    Id int `gorm:"primaryKey"`
    Title string
    Phone string `gorm:"not null;default:null"`
    Remark string `gorm:"not null;default:null"`
    MenuItems []MenuItem
}

type MenuItem struct {
    Id int `gorm:"primaryKey"`
    MenuID int
    Name string
    Price int `gorm:"not null;default:null"`
    Remark string `gorm:"not null;default:null"`
}
