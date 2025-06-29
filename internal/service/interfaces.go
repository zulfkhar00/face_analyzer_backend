package service

import (
	"context"
	"mime/multipart"

	"github.com/zulfkhar00/cosmetics-backend/internal/domain"
	"github.com/zulfkhar00/cosmetics-backend/internal/handler/dto"
)

// ProductService defines the business logic operations for products.
type ProductService interface {
	GetProductByBarcode(ctx context.Context, barcode string) (*domain.Product, error)
}

type FaceAnalysisService interface {
	AnalyzeAndUploadFaceImage(ctx context.Context, uid string, faceImage *multipart.FileHeader) error
	GetUserFaceCondition(ctx context.Context, uid string) (*domain.Face, error)
	AddProductToRoutine(ctx context.Context, uid string, productID string, routineType dto.RoutineType) error
}
