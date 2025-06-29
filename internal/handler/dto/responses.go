package dto

type GetProductResponse struct {
	ID                  string   `json:"id,omitempty"` // omitempty if ID might be empty for external products
	Barcode             string   `json:"barcode"`
	ProductName         string   `json:"productName"`
	Brand               string   `json:"brand"`
	ProductQuantity     float64  `json:"productQuantity"`
	ProductQuantityUnit string   `json:"productQuantityUnit"`
	Ingredients         []string `json:"ingredients"`
	Source              string   `json:"source"`
}

type SendUserFaceImageResponse struct {
	Message string `json:"message,omitempty"`
	Error   error  `json:"error,omitempty"`
}

type GetUserFaceConditionResponse struct {
	FaceCondition FaceCondition `json:"faceCondition,omitempty"`
	Error         error         `json:"error,omitempty"`
}

type FaceCondition struct {
	Probabilities    map[string]float64 `json:"probabilities"`
	OverallScore     float32            `json:"overallScore"`
	OverallCondition string             `json:"overallCondition"`
}

type AddProductToRoutineResonse struct {
	Message string `json:"message,omitempty"`
	Error   error  `json:"error,omitempty"`
}
