package handler

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/zulfkhar00/cosmetics-backend/internal/handler/dto"
	"github.com/zulfkhar00/cosmetics-backend/internal/service"
)

type FaceAnalysisHandler struct {
	faceAnalysisService service.FaceAnalysisService
}

func NewFaceAnalysisHandler(svc service.FaceAnalysisService) *FaceAnalysisHandler {
	return &FaceAnalysisHandler{
		faceAnalysisService: svc,
	}
}

func (h *FaceAnalysisHandler) SendUserFaceImage(ctx context.Context, c *app.RequestContext) {
	var req dto.SendUserFaceImageRequest
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{"error": "invalid request data"})
		return
	}

	err := h.faceAnalysisService.AnalyzeAndUploadFaceImage(ctx, req.UID, req.FaceImage)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, dto.SendUserFaceImageResponse{Error: err})
		return
	}

	c.JSON(consts.StatusOK, dto.SendUserFaceImageResponse{Message: "Face image processed successfully"})
}

func (h *FaceAnalysisHandler) GetUserFaceCondition(ctx context.Context, c *app.RequestContext) {
	var req dto.GetUserFaceConditionRequest
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(consts.StatusBadRequest, dto.GetUserFaceConditionResponse{Error: errors.New("invalid request data")})
		return
	}

	face, err := h.faceAnalysisService.GetUserFaceCondition(ctx, req.UID)
	if err != nil {
		c.JSON(consts.StatusBadRequest, dto.GetUserFaceConditionResponse{Error: err})
		return
	}

	resp := dto.GetUserFaceConditionResponse{
		FaceCondition: dto.FaceCondition{
			Probabilities:    face.Probabilities,
			OverallScore:     face.OverallScore,
			OverallCondition: face.OverallCondition,
		},
	}

	c.JSON(consts.StatusOK, resp)
}
