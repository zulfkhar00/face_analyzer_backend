package handler

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	dto "github.com/zulfkhar00/cosmetics-backend/internal/handler/dto/response"
	"github.com/zulfkhar00/cosmetics-backend/internal/service"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(svc service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: svc,
	}
}

func (h *ProductHandler) GetProductByBarcode(ctx context.Context, c *app.RequestContext) {
	barcode := c.Param("barcode")

	product, err := h.productService.GetProductByBarcode(ctx, barcode)

	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			c.JSON(consts.StatusNotFound, map[string]string{"error": fmt.Sprintf("Product with barcode '%s' not found.", barcode)})
			return
		}

		log.Printf("Handler: Error in GetProductByBarcode service call for barcode '%s': %v", barcode, err)
		c.JSON(consts.StatusInternalServerError, map[string]string{"error": "Internal server error fetching product."})
		return
	}

	responseProduct := dto.GetProductResponse{
		ID:                  product.ID,
		Barcode:             product.Barcode,
		ProductName:         product.ProductName,
		Brand:               product.Brand,
		ProductQuantity:     product.ProductQuantity,
		ProductQuantityUnit: product.ProductQuantityUnit,
		Ingredients:         product.Ingredients,
		Source:              product.Source,
	}

	c.JSON(consts.StatusOK, responseProduct)
}
