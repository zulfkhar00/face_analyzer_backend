package domain

// ProductRepository defines operations for product data access
type ProductRepository interface {
	GetByBarcode(barcode string) (*Product, error)
}

type FaceRepository interface {
	UploadFaceImage(face *Face) error
	GetUserFaceCondition(uid string) (*Face, error)
	AddProductToRoutine(uid string, productID string, routineType string) error
}
