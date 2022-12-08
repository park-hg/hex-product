package api

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"product_catalog/domain"
)

type handler struct {
	productService domain.Service
}

func NewHandler(productService domain.Service) ProductHandler {
	return &handler{productService: productService}
}

func (h *handler) Get(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	code := ctx.Params("code")
	p, err := h.productService.Find(code)
	if err != nil {
		return err
	}
	return ctx.JSON(p)
}

func (h *handler) Post(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	p := &domain.Product{}
	err := ctx.BodyParser(p)
	if err != nil {
		return err
	}

	err = h.productService.Store(p)
	if err != nil {
		return err
	}
	return ctx.Status(http.StatusCreated).JSON(p)
}

func (h *handler) Put(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	var p domain.Product
	err := ctx.BodyParser(p)
	if err != nil {
		return err
	}

	err = h.productService.Update(&p)
	if err != nil {
		return err
	}
	return ctx.Status(http.StatusNoContent).JSON(&p)
}

func (h *handler) Delete(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	var p domain.Product
	err := ctx.BodyParser(p)
	if err != nil {
		return err
	}

	err = h.productService.Delete(p.Code)
	if err != nil {
		return err
	}
	return ctx.Status(http.StatusNoContent).JSON(&p)
}

func (h *handler) GetAll(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	ps, err := h.productService.FindAll()
	if err != nil {
		return err
	}
	return ctx.Status(http.StatusOK).JSON(ps)
}
