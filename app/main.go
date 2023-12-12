package main

import (
	"context"
	"log"

	"github.com/ecommerce/config"
	"github.com/ecommerce/domain/auth"
	"github.com/ecommerce/domain/category"
	"github.com/ecommerce/domain/product"
	"github.com/ecommerce/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	app := fiber.New(fiber.Config{
		AppName: "Ecommerce Services",
		Prefork: true,
	})

	app.Use(logger.New())

	err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Println("error when try to LoadConfig with error :", err.Error())
	}

	db, err := database.ConnectSQLXPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	rdb, err := database.ConnectRedis(context.Background(), config.Cfg.Redis)
	if err != nil {
		panic(err)
	}

	auth.RegisterServiceAuth(app, auth.DB{Dbx: db, Redis: rdb, Cfg: config.Cfg.JWT})
	category.RegisterServiceCategory(app, category.DB{Dbx: db})
	product.RegisterServiceProduct(app, product.DB{Dbx: db})

	app.Listen(config.Cfg.App.Port)
}
