package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/zulfkhar00/cosmetics-backend/internal/domain"
	"github.com/zulfkhar00/cosmetics-backend/pkg/apiclient"
)

type productService struct {
	productRepo           domain.ProductRepository
	openBeautyFactsClient apiclient.OpenBeautyFactsClient
}

func NewProductService(
	productRepo domain.ProductRepository,
	openBeautyFactsClient apiclient.OpenBeautyFactsClient,
) ProductService {
	return &productService{
		productRepo:           productRepo,
		openBeautyFactsClient: openBeautyFactsClient,
	}
}

func (s *productService) GetProductByBarcode(ctx context.Context, barcode string) (*domain.Product, error) {
	product, err := s.productRepo.GetByBarcode(barcode)
	if err == nil {
		return product, nil
	}

	if !errors.Is(err, domain.ErrNotFound) {
		log.Printf("Service: Database error fetching product '%s': %v", barcode, err)
		return nil, fmt.Errorf("failed to retrieve product from internal database: %w", err)
	}

	openBeautyFactsResp, err := s.openBeautyFactsClient.GetProductByBarcode(ctx, barcode)
	if err != nil {
		if errors.Is(err, apiclient.ErrOpenBeautyFactsNotFound) {
			log.Printf("Service: Product '%s' not found in OpenBeautyFacts API.", barcode)
			return nil, fmt.Errorf("%w: %s", ErrNotFound, "product not found in any source")
		}
		log.Printf("Service: Error fetching product '%s' from OpenBeautyFacts API: %v", barcode, err)
		return nil, fmt.Errorf("failed to retrieve product from external API: %w", err)
	}

	ingredients := make([]string, 0, len(openBeautyFactsResp.Product.Ingredients))
	for _, ingredient := range openBeautyFactsResp.Product.Ingredients {
		ingredients = append(ingredients, ingredient.Text)
	}

	productFromExternalAPI := &domain.Product{
		ID:                  "", // No internal ID for externally fetched product yet
		Barcode:             openBeautyFactsResp.Product.Code,
		ProductName:         openBeautyFactsResp.Product.ProductNameEn,
		Brand:               openBeautyFactsResp.Product.Brands,
		ProductQuantityUnit: openBeautyFactsResp.Product.ProductQuantityUnit,
		ProductQuantity:     openBeautyFactsResp.Product.ProductQuantity,
		Ingredients:         ingredients,
		Source:              "open_beauty_facts",
	}

	return productFromExternalAPI, nil
}
