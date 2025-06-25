package dto

type GetProductRequest struct {
	Barcode string `json:"barcode" binding:"required"`
}
