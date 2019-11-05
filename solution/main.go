package main

import (
	"database/sql"

	"detectify/solution/imagehandler/repository"
	"detectify/solution/imagehandler/usecase"
	"detectify/solution/imagehttphandler"
	"detectify/solution/middleware"
	"detectify/solution/pkg/helper"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	pvError := viper.ReadInConfig()
	if pvError != nil {
		panic(pvError)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {

	var imgHelper helper.ImageHelper
	imgHelper.APIUser = viper.GetString("folderConfig.APIUser")
	imgHelper.APIkey = viper.GetString("folderConfig.APIKey")
	imgHelper.Path = viper.GetString("folderConfig.path")

	pvDbserverName := viper.GetString(`database.server`)
	pvDbPort := viper.GetString(`database.port`)
	pvDbName := viper.GetString(`database.databaseName`)

	pvUserName := viper.GetString(`database.userName`)
	pvPassword := viper.GetString(`database.password`)

	pvConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		pvUserName, pvPassword, pvDbserverName, pvDbPort, pvDbName)

	pvDbConnection, pvError := sql.Open("mysql", pvConnString)
	if pvError != nil && viper.GetBool("debug") {
		fmt.Println(pvError)
	}
	defer pvDbConnection.Close()

	pvError = pvDbConnection.Ping()

	if pvError != nil {
		log.Fatal(pvError)
		os.Exit(1)
	}

	defer func() {
		pvError := pvDbConnection.Close()
		if pvError != nil {
			log.Fatal(pvError)
		}
	}()

	pvEcho := echo.New()
	pvMiddle := middleware.InitMiddleware()
	pvEcho.Use(pvMiddle.CORS)

	pvImageRepo := repository.NewImageRepository(pvDbConnection)
	pvImageHelper := imgHelper.NewImageHelper(&imgHelper)
	PVTimeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	pvImageBusiness := usecase.NewImageBusiness(pvImageRepo, *pvImageHelper, PVTimeoutContext)
	imagehttphandler.NewImageHandler(pvEcho, pvImageBusiness)
	log.Fatal(pvEcho.Start(viper.GetString("server.address")))
}
