package models

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) (string, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return "", err
	}

	return os.Getenv(key), nil
}

var DB *gorm.DB

func ConnectToBDD() {
	userDb, err := goDotEnvVariable("USERDB")
	if err != nil {
		userDb = "root"
	}

	passDb, err := goDotEnvVariable("PASSDB")
	if err != nil {
		passDb = "root"
	}

	nameDb, err := goDotEnvVariable("NAMEDB")
	if err != nil {
		nameDb = "cqrs"
	}

	DB, err = gorm.Open("mysql", userDb+":"+passDb+"@(db)/"+nameDb+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
	}
}
