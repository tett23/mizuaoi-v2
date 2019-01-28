package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DBConfig DBConfig
type DBConfig struct {
	User     string
	Password string
	DBName   string
}

func dbConfig() *DBConfig {
	var user = os.Getenv("DB_USER")
	if user == "" {
		user = "root"
	}
	var password = os.Getenv("DB_PASSWORD")
	if password == "" {
		password = ""
	}
	var name = os.Getenv("DB_NAME")
	if name == "" {
		name = "mizuaoi-v2_development"
	}

	return &DBConfig{
		User:     user,
		Password: password,
		DBName:   name,
	}
}

func main() {
	var config = dbConfig()
	var connectionConfig = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=UTC", config.User, config.Password, config.DBName)

	db, err := gorm.Open("mysql", connectionConfig)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
