package handlers

import (
	"github.com/Montheankul-K/go-redis/services"
	"github.com/gofiber/fiber/v2"
)

type catalogHandler struct {
	catalogService services.CatalogService
}

func NewCatalogHandler(catalogService services.CatalogService) CatalogHandler {
	return catalogHandler{catalogService: catalogService}
}

func (h catalogHandler) GetProducts(c *fiber.Ctx) error {
	products, err := h.catalogService.GetProducts()
	if err != nil {
		return err
	}

	response := fiber.Map{
		"status":   "ok",
		"products": products,
	}

	return c.JSON(response)
}
