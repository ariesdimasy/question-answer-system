package main

import (
	_dbDriver "acp-final/drivers/databases"

	"log"
	"time"

	_userUsecase "acp-final/business/users"
	_userController "acp-final/controllers/users"
	_userRepo "acp-final/drivers/databases/users"

	_productUseCase "acp-final/business/products"
	_productController "acp-final/controllers/products"
	_productRepo "acp-final/drivers/databases/products"

	_route "acp-final/apps/routes"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("apps/configs/config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	db := configDB.ConnectDB()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepo := _userRepo.NewUserRepository(db)
	userUseCase := _userUsecase.NewUserUsecase(userRepo, timeoutContext)
	userController := _userController.NewUserController(userUseCase)

	productRepo := _productRepo.NewUserRepository(db)
	productUseCase := _productUseCase.NewProductUsecase(productRepo, timeoutContext)
	productController := _productController.NewProductController(productUseCase)

	e := echo.New()
	routesInit := _route.ControllerList{
		UserController:    *userController,
		ProductController: *productController,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
