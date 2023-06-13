package main

import (
	"synapsis-backend/configs"
	_ "synapsis-backend/docs"
	"synapsis-backend/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           Synapsis Online Store API Documentation
// @version         1.0
// @termsOfService  http://swagger.io/terms/

// @contact.name   Daniel Capah
// @contact.url    https://github.com/capahdan

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// https://rest-api-7qon5jxieq-et.a.run.app/
// https://rest-api-7qon5jxieq-et.a.run.app/

// @host      rest-api-7qon5jxieq-et.a.run.app
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http:localhost:8080", "https:rest-api-7qon5jxieq-et.a.run.app"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	db, err := configs.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = configs.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	routes.Init(e, db)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
