package src

import (
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

func UserRegister(c *gin.Context) {
    user:=UserRequest{}
    err:=c.BindJSON(&user)
    if err!=nil {
        c.String(400, "param error %s", err.Error())
        return
    }

    if CheckAccountDuplicate(user.Account) {
        c.String(409, "account duplicated")
        return
    }

    password,err:=HashPassword(user.Password)
    if err != nil {
        c.String(400, "password error %s", err.Error())
        return
    }

    userCreate:=User {
        Name: user.Name,
        Account: user.Account,
        Password: password}

    if err:=DBHelper.Create(&userCreate).Error; err != nil {
        c.String(400, "failing create %s", err.Error())
    } else {
        c.String(200, "success")
    }
}

func CheckAccountDuplicate(account string) bool {
    user:=User{}
    err:=DBHelper.Where("account = ?", account).First(&user).Error

    return err==nil
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
