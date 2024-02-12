package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	model "github.com/ValSpp/gin-gorm-restapi/models"
)

var DB *gorm.DB
var err error

const DNS = "postgresql://postgresql:noval123@127.0.0.1:5433/kubikitapidb?sslmode=require"

func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		fmt.Println(err.Error())
		panic("tidak dapat terhubung ke database")
	}
	DB.AutoMigrate(&model.Users{})
}
