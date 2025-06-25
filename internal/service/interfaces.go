package service

import (
	"context"

	"github.com/zulfkhar00/cosmetics-backend/internal/domain"
)

// ProductService defines the business logic operations for products.
type ProductService interface {
	GetProductByBarcode(ctx context.Context, barcode string) (*domain.Product, error)
}
