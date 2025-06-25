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
