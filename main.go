package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"product_catalog/api"
	"product_catalog/domain"
	"product_catalog/repository"
)

func main() {

	repo := repository.NewMysqlRepository()
	service := domain.NewProductService(repo)
	handler := api.NewHandler(service)

	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello world!")
	})

	app.Get("/products/{code}", handler.Get)
	app.Post("/products/{code}", handler.Post)
	app.Put("/products/{code}", handler.Put)
	app.Delete("/products/{code}", handler.Delete)
	app.Get("/products", handler.GetAll)

	log.Fatal(app.Listen(":3000"))
}
