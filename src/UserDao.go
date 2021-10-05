package src

import (
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "time"
    "math/rand"
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

func FindUser(id uint) (user User) {
    DBHelper.Where("id = ?", id).First(&user)
    return
}

func GetUserByAccount(account string) (user User) {
    DBHelper.Where("account = ?", account).First(&user)
    return
}


func UserLogin(c *gin.Context) {
    request := map[string]string{}
    err:=c.BindJSON(&request)
    if err!=nil {
        c.String(400, "binding error %s", err.Error())
        return
    }

    user := GetUserByAccount(request["account"])
    if passCheck := CheckPasswordHash(request["password"], user.Password); passCheck == false {
        c.String(401, "wrong password  %s", user.Password)
        return
    }

    GenerateToken(c, user)
}

func GetRefreshToken(user User) string {
    if (user.RefreshToken == "") {
        user.RefreshToken = GenerateRefreshToken();
        DBHelper.Save(&user)
    }

    return user.RefreshToken
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateRefreshToken() (string) {
    rand.Seed(time.Now().UnixNano())
    b := make([]rune, 10)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}
