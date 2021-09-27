package src

type UserRequest struct {
    Name string `json:"name" binding:"required"`
    Account string `json:"account" binding:"required"`
    Password string `json:"password" binding:"required,min=6"`
}

type User struct {
    Id uint `gorm:"primaryKey"`
    Name string
    Account string
    Password string
}
