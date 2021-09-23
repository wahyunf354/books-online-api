package main

import (
	_usersdb "books_online_api/drivers/databases/users"
	_mysqlDriver "books_online_api/drivers/mysql"
	"log"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Server RUN on Debug mode")
	}
}

func DbMigration(db *gorm.DB) {
	db.AutoMigrate(&_usersdb.Users{})
}

func main() {
	// init connection database
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
}
