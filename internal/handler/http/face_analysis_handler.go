package handler

import (
	"context"

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
		c.JSON(consts.StatusBadRequest, map[string]string{"error": "Invalid request data"})
		return
	}

	err := h.faceAnalysisService.AnalyzeAndUploadFaceImage(ctx, req.UID, req.FaceImage)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, map[string]string{"error": "Failed to process face image"})
		return
	}

	c.JSON(consts.StatusOK, map[string]string{"message": "Face image processed successfully"})
}
