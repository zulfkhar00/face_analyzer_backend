package domain

// Product represents a skincare product
type Product struct {
	ID                  string   // Corresponds to p.id AS product_id (UUID from DB)
	Barcode             string   // Corresponds to p.code AS barcode (string from DB)
	ProductName         string   // Corresponds to p.product_name
	Brand               string   // Corresponds to p.brands AS brand
	ProductQuantity     float64  // Corresponds to p.product_quantity
	ProductQuantityUnit string   // Corresponds to p.product_quantity_unit
	Ingredients         []string // Corresponds to json_agg(DISTINCT i.ingredient_name) AS ingredients
	Source              string   // This field is application-specific, not from DB, so will be set by service/handler
}
