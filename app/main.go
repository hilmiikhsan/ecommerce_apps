package main

import (
	"log"

	"github.com/ecommerce/config"
	"github.com/ecommerce/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	router := fiber.New(fiber.Config{
		AppName: "Ecommerce Services",
		Prefork: true,
	})

	err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Println("error when try to LoadConfig with error :", err.Error())
	}

	_, err = database.ConnectSQLXPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	router.Listen(config.Cfg.App.Port)
}
