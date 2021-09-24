package main

import (
	"books_online_api/app/middlewares"
	"books_online_api/app/routes"
	_userUseCase "books_online_api/business/users"
	_userController "books_online_api/controllers/users"
	_userRepository "books_online_api/drivers/databases/users"
	_userdb "books_online_api/drivers/databases/users"
	_mysqlDriver "books_online_api/drivers/mysql"
	"log"
	"time"

	"github.com/labstack/echo/v4"
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
	err := db.AutoMigrate(&_userdb.Users{})
	if err != nil {
		panic(err)
	}
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

	configJwt := middlewares.ConfigJwt{
		SecretJwt: viper.GetString("jwt.secret"),
		ExpiredDuration: viper.GetInt("jwt.expired"),
	}

	Conn := configDB.InitialDB()
	DbMigration(Conn)

	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepository := _userRepository.NewMysqlUserRepository(Conn)
	userUseCase := _userUseCase.NewUserUseCase(userRepository, timeoutContext, configJwt)
	userController := _userController.NewUserController(userUseCase)

	routesInit := routes.ControllerList{
		JWTMiddleware: configJwt.Init(),
		UserController: *userController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))

}