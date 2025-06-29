package service

import (
	"context"
	"mime/multipart"

	"github.com/zulfkhar00/cosmetics-backend/internal/domain"
	"github.com/zulfkhar00/cosmetics-backend/internal/utility"
	"github.com/zulfkhar00/cosmetics-backend/scripts"
)

type faceAnalysisService struct {
	productRepo domain.FaceRepository
}

func NewFaceAnalysisService(productRepo domain.FaceRepository) FaceAnalysisService {
	return &faceAnalysisService{
		productRepo: productRepo,
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

	return s.productRepo.UploadFaceImage(face)
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
