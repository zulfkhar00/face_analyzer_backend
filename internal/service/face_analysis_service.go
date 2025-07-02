package service

import (
	"context"
	"mime/multipart"

	"github.com/zulfkhar00/cosmetics-backend/internal/domain"
	"github.com/zulfkhar00/cosmetics-backend/internal/handler/dto"
	"github.com/zulfkhar00/cosmetics-backend/internal/utility"
	"github.com/zulfkhar00/cosmetics-backend/scripts"
)

type faceAnalysisService struct {
	faceRepo domain.FaceRepository
}

func NewFaceAnalysisService(faceRepo domain.FaceRepository) FaceAnalysisService {
	return &faceAnalysisService{
		faceRepo: faceRepo,
	}
}

func (s *faceAnalysisService) AnalyzeAndUploadFaceImage(ctx context.Context, uid string, faceImage *multipart.FileHeader) error {
	tempPath := "/tmp/uploaded_face_images/" + faceImage.Filename
	if err := utility.SaveUploadedFile(faceImage, tempPath); err != nil {
		return err
	}

	_, probs, err := scripts.RunFaceAnalysisPython(tempPath)
	if err != nil {
		return err
	}

	overallScore, overallCondition := s.GetFaceHealthScore(ctx, probs)

	face := &domain.Face{
		UID:              uid,
		Probabilities:    probs,
		OverallScore:     overallScore,
		OverallCondition: overallCondition,
	}

	return s.faceRepo.UploadFaceImage(face)
}

func (s *faceAnalysisService) GetFaceHealthScore(ctx context.Context, predictions map[string]float64) (score float32, overallCondition string) {
	var total float64
	for _, prob := range predictions {
		total += prob
	}
	n := float64(len(predictions))
	avg := total / n
	score = float32(100 * (1 - avg))

	if score >= 90 {
		overallCondition = "Excellent"
	} else if score >= 75 {
		overallCondition = "Good"
	} else if score >= 60 {
		overallCondition = "Moderate"
	} else if score >= 40 {
		overallCondition = "Poor"
	} else {
		overallCondition = "Very Poor"
	}
	return
}

func (s *faceAnalysisService) GetUserFaceCondition(ctx context.Context, uid string) (*domain.Face, error) {
	face, err := s.faceRepo.GetUserFaceCondition(uid)
	if err != nil {
		return nil, err
	}

	return face, nil
}

func (s *faceAnalysisService) AddProductToRoutine(ctx context.Context, uid string, productID string, routineType dto.RoutineType) error {
	err := s.faceRepo.AddProductToRoutine(uid, productID, string(routineType))
	if err != nil {
		return err
	}

	return nil
}

func (s *faceAnalysisService) GetRoutines(ctx context.Context, uid string) ([]*domain.UserRoutineProduct, error) {
	routineProducts, err := s.faceRepo.GetRoutines(ctx, uid)
	if err != nil {
		return nil, err
	}

	for _, routine := range routineProducts {
		for _, product := range routine.Products {
			ingredients := make([]string, 0)
			for _, ingredient := range product.IngredientsWithImpact {
				ingredients = append(ingredients, ingredient.IngredientName)
			}
			product.Ingredients = ingredients
		}
	}

	return routineProducts, nil
}
