package dto

import "mime/multipart"

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
