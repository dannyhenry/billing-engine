package main

import (
	"fmt"

	"github.com/dannyhenry/billing/common/application"
	"github.com/dannyhenry/billing/common/config"
	"github.com/dannyhenry/billing/common/database"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	_ = godotenv.Load()
	app := application.App{
		E:         echo.New(),
		DBManager: database.NewDatabaseManager(),
	}

	dsn := database.PostgresURI(
		config.GetConfig().DbUser,
		config.GetConfig().DbPassword,
		fmt.Sprintf(`%s:%s`, config.GetConfig().DbHost, config.GetConfig().DbPort),
		config.GetConfig().DbName,
		config.GetConfig().DBSchema)

	app.InitializeDatabase(
		dsn,
		config.GetConfig().DbMaxIdleConns,
		config.GetConfig().DbMaxOpenConns,
	)
	app.Initialize()

	app.Start(":" + config.GetConfig().AppPort)
}
