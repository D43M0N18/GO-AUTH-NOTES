package utils

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "os"
    "time"
    "go-auth-notes/models"
)

var DB *gorm.DB

func ConnectDB() error {
    dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(mysql:3306)/" + os.Getenv("DB_NAME") + "?parseTime=true"
    
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }

    sqlDB, _ := DB.DB()
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)
    sqlDB.SetConnMaxLifetime(time.Hour)

    return DB.AutoMigrate(&models.User{}, &models.Note{})
}
