package client_db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"group_shopping_mall/utils/utils"
)

var (
	clientDB *gorm.DB
)

type DBConfig struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
	Database string `json:"database,omitempty"`
}

func InitDB(path string) {
	config := utils.GetConfig[DBConfig](path)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	clientDB = db
}

func GetDB() *gorm.DB {
	return clientDB
}
