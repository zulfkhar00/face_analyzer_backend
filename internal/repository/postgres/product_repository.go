package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/zulfkhar00/cosmetics-backend/internal/domain"
)

type PostgresProductRepository struct {
	db *sql.DB // connection pool
}

func NewProductRepository(db *sql.DB) *PostgresProductRepository {
	return &PostgresProductRepository{db: db}
}

func (r *PostgresProductRepository) GetByBarcode(barcode string) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Add context with timeout
	defer cancel()

	query := `
        SELECT
            p.id AS product_id,
            p.code AS barcode,
            p.product_name,
            p.brands AS brand,
            p.product_quantity,
            p.product_quantity_unit,
            json_agg(DISTINCT i.ingredient_name) FILTER (WHERE i.ingredient_name IS NOT NULL) AS ingredients
        FROM
            products p
        LEFT JOIN
            product_ingredients pi ON pi.product_id = p.id
        LEFT JOIN
            ingredients i ON i.id = pi.ingredient_id
        WHERE
            p.code = $1
        GROUP BY
            p.id;
    `

	product := &domain.Product{}
	var ingredientsJSON []byte

	err := r.db.QueryRowContext(ctx, query, barcode).Scan(
		&product.ID,                  // p.id AS product_id
		&product.Barcode,             // p.code AS barcode
		&product.ProductName,         // p.product_name
		&product.Brand,               // p.brands AS brand
		&product.ProductQuantity,     // p.product_quantity
		&product.ProductQuantityUnit, // p.product_quantity_unit
		&ingredientsJSON,             // json_agg(...) AS ingredients
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product with barcode %s not found: %w", barcode, domain.ErrNotFound)
		}
		return nil, fmt.Errorf("failed to get product by barcode %s: %w", barcode, err)
	}

	if len(ingredientsJSON) > 0 {
		if err := json.Unmarshal(ingredientsJSON, &product.Ingredients); err != nil {
			return nil, fmt.Errorf("failed to unmarshal ingredients JSON for barcode %s: %w", barcode, err)
		}
	} else {
		product.Ingredients = []string{}
	}

	product.Source = "internal"

	return product, nil
}
