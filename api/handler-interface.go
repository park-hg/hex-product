package api

import (
	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	Get(ctx *fiber.Ctx) error
	Post(ctx *fiber.Ctx) error
	Put(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
}
