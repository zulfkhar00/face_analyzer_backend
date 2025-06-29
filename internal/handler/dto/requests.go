package dto

import "mime/multipart"

type RoutineType string

const (
	MorningRoutine RoutineType = "morning"
	EveningRoutine RoutineType = "evening"
	BothRoutines   RoutineType = "both"
)

func (rt RoutineType) IsValid() bool {
	return rt == MorningRoutine || rt == EveningRoutine || rt == BothRoutines
}

type GetProductRequest struct {
	Barcode string `json:"barcode" binding:"required"`
}

type SendUserFaceImageRequest struct {
	UID       string                `form:"uid" binding:"required"`
	FaceImage *multipart.FileHeader `form:"image" binding:"required"`
}

type GetUserFaceConditionRequest struct {
	UID string `json:"uid" binding:"required"`
}

type AddProductToRoutineRequest struct {
	UID         string      `json:"uid" binding:"required"`
	ProductID   string      `json:"productID" binding:"required"`
	RoutineType RoutineType `json:"routineType" binding:"required,oneof=morning evening both"`
}
