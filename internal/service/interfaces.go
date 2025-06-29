package service

import (
	"context"
	"mime/multipart"

	"github.com/zulfkhar00/cosmetics-backend/internal/domain"
)

// ProductService defines the business logic operations for products.
type ProductService interface {
	GetProductByBarcode(ctx context.Context, barcode string) (*domain.Product, error)
}

type FaceAnalysisService interface {
	AnalyzeAndUploadFaceImage(ctx context.Context, uid string, faceImage *multipart.FileHeader) error
	GetUserFaceCondition(ctx context.Context, uid string) (*domain.Face, error)
}
