package src

import (
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
    "github.com/gin-gonic/gin"
    "net/http"
)

const KEY = "owf_session_key"

func EnableCookieSession() gin.HandlerFunc {
    store := cookie.NewStore([]byte(KEY))
    return sessions.Sessions("owf_session", store)
}

func AuthSessionMiddle() gin.HandlerFunc {
    return func(c *gin.Context) {
        session := sessions.Default(c)
        sessionValue := session.Get("userId")
        if sessionValue == nil {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Unauthorized",
            })
            c.Abort()
            return
        }

        c.Set("userId", sessionValue.(uint))
        c.Next()
        return
    }
}

func SaveAuthSession(c *gin.Context, id uint) {
    session := sessions.Default(c)
    session.Set("userId", id)
    session.Save()
}

func ClearAuthSession(c *gin.Context) {
    session := sessions.Default(c)
    session.Clear()
    session.Save()
}

func HasLoginSession(c *gin.Context) bool {
    session := sessions.Default(c)
    if sessionValue := session.Get("userId"); sessionValue == nil {
        return false
    }
    return true
}

func GetSessionUserId(c *gin.Context) uint {
    session := sessions.Default(c)
    sessionValue := session.Get("userId")
    if sessionValue == nil {
        return 0
    }
    return sessionValue.(uint)
}

func GetUserSession(c *gin.Context) map[string]interface{} {
    hasSession := HasLoginSession(c)
    userName := ""
    if hasSession {
        userId := GetSessionUserId(c)
        userName = FindUser(userId).Name
    }
    data := make(map[string]interface{})
    data["hasSession"] = hasSession
    data["userName"] = userName
    return data
}
