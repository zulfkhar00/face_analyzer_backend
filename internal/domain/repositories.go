package domain

// ProductRepository defines operations for product data access
type ProductRepository interface {
	GetByBarcode(barcode string) (*Product, error)
}
