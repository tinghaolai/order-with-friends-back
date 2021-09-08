package src

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "time"
)

var DBHelper *gorm.DB
var err error
func InitDB() {
    dsn := "root:root@tcp(127.0.0.1:3307)/owf?charset=utf8&parseTime=True&loc=Local"
    DBHelper, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })

    if err != nil {
        ShutDownServer(err)
        return
    }

    sqlDB, _ := DBHelper.DB()
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)
    sqlDB.SetConnMaxLifetime(time.Hour)
}
