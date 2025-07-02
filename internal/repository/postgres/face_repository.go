package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/zulfkhar00/cosmetics-backend/internal/domain"
)

type PostgresFaceRepository struct {
	db *sql.DB // connection pool
}

func NewFaceRepository(db *sql.DB) *PostgresFaceRepository {
	return &PostgresFaceRepository{db: db}
}

func (r *PostgresFaceRepository) UploadFaceImage(face *domain.Face) error {
	probsJSON, err := json.Marshal(face.Probabilities)
	if err != nil {
		return fmt.Errorf("failed to marshal probabilities: %w", err)
	}

	query := `
		INSERT INTO user_face_condition (uid, probabilities, overall_score, overall_condition)
		VALUES ($1, $2, $3, $4)
	`

	_, err = r.db.Exec(query, face.UID, probsJSON, face.OverallScore, face.OverallCondition)
	if err != nil {
		return fmt.Errorf("failed to insert face image data: %w", err)
	}

	return nil
}

func (r *PostgresFaceRepository) GetUserFaceCondition(uid string) (*domain.Face, error) {
	query := `
		SELECT probabilities, overall_score, overall_condition
		FROM user_face_condition
		WHERE uid = $1
		ORDER BY created_at DESC
		LIMIT 1
	`

	row := r.db.QueryRow(query, uid)

	var (
		probJSON         []byte
		overallScore     float64
		overallCondition string
	)

	if err := row.Scan(&probJSON, &overallScore, &overallCondition); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no record found for uid: %s", uid)
		}
		return nil, fmt.Errorf("failed to scan face condition: %w", err)
	}

	var probabilities map[string]float64
	if err := json.Unmarshal(probJSON, &probabilities); err != nil {
		return nil, fmt.Errorf("failed to unmarshal probabilities: %w", err)
	}

	face := &domain.Face{
		UID:              uid,
		Probabilities:    probabilities,
		OverallScore:     float32(overallScore),
		OverallCondition: overallCondition,
	}

	return face, nil
}

func (r *PostgresFaceRepository) AddProductToRoutine(uid string, productID string, routineType string) error {
	query := `
		INSERT INTO user_skincare_routine (uid, product_id, routine_type)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(query, uid, productID, routineType)
	if err != nil {
		return fmt.Errorf("failed to insert skincare routine: %w", err)
	}

	return nil
}

func (r *PostgresFaceRepository) GetRoutines(ctx context.Context, uid string) ([]*domain.UserRoutineProduct, error) {
	query := `
		SELECT
			usr.routine_type,
			usr.created_at as routine_created_at,
			p.id::text,
			COALESCE(p.code, '') AS code,
			COALESCE(p.product_name, '') AS product_name,
			COALESCE(p.brands, '') AS brand,
			COALESCE(p.product_quantity, 0.0) AS product_quantity,
			COALESCE(p.product_quantity_unit, '') AS product_quantity_unit,
			COALESCE(
				(
					SELECT json_agg(
						json_build_object(
							'ingredient_name', i.ingredient_name,
							'blackheads_impact', ii.blackheads_impact,
							'blackheads_confidence', ii.blackheads_confidence,
							'dark_circles_on_face_impact', ii.dark_circles_on_face_impact,
							'dark_circles_on_face_confidence', ii.dark_circles_on_face_confidence,
							'dark_spots_on_face_impact', ii.dark_spots_on_face_impact,
							'dark_spots_on_face_confidence', ii.dark_spots_on_face_confidence,
							'dry_skin_impact', ii.dry_skin_impact,
							'dry_skin_confidence', ii.dry_skin_confidence,
							'dull_skin_impact', ii.dull_skin_impact,
							'dull_skin_confidence', ii.dull_skin_confidence,
							'eye_bags_impact', ii.eye_bags_impact,
							'eye_bags_confidence', ii.eye_bags_confidence,
							'face_redness_impact', ii.face_redness_impact,
							'face_redness_confidence', ii.face_redness_confidence,
							'forehead_wrinkles_impact', ii.forehead_wrinkles_impact,
							'forehead_wrinkles_confidence', ii.forehead_wrinkles_confidence,
							'hormonal_acne_impact', ii.hormonal_acne_impact,
							'hormonal_acne_confidence', ii.hormonal_acne_confidence,
							'large_pores_on_face_impact', ii.large_pores_on_face_impact,
							'large_pores_on_face_confidence', ii.large_pores_on_face_confidence,
							'oily_skin_impact', ii.oily_skin_impact,
							'oily_skin_confidence', ii.oily_skin_confidence,
							'razor_bumps_impact', ii.razor_bumps_impact,
							'razor_bumps_confidence', ii.razor_bumps_confidence,
							'rough_texture_on_face_impact', ii.rough_texture_on_face_impact,
							'rough_texture_on_face_confidence', ii.rough_texture_on_face_confidence,
							'sebaceous_filaments_impact', ii.sebaceous_filaments_impact,
							'sebaceous_filaments_confidence', ii.sebaceous_filaments_confidence,
							'under_eye_wrinkles_impact', ii.under_eye_wrinkles_impact,
							'under_eye_wrinkles_confidence', ii.under_eye_wrinkles_confidence
						)
					)
					FROM product_ingredients pi
					JOIN ingredients i ON pi.ingredient_id = i.id
					LEFT JOIN ingredient_impact ii ON i.id = ii.ingredient_id
					WHERE pi.product_id = p.id
				),
				'[]'::json
			) AS ingredients_with_impact
		FROM
			user_skincare_routine usr
		JOIN
			products p ON usr.product_id = p.id
		WHERE
			usr.uid = $1
		ORDER BY
			usr.routine_type, routine_created_at DESC;
	`

	rows, err := r.db.QueryContext(ctx, query, uid)
	if err != nil {
		return nil, fmt.Errorf("failed to query user routine products: %w", err)
	}
	defer rows.Close()

	routineMap := make(map[string]*domain.UserRoutineProduct)

	for rows.Next() {
		var routineType string
		var routineCreatedAt time.Time
		var product domain.Product
		var ingredientsJSON []byte

		if err := rows.Scan(
			&routineType,
			&routineCreatedAt,
			&product.ID,
			&product.Barcode,
			&product.ProductName,
			&product.Brand,
			&product.ProductQuantity,
			&product.ProductQuantityUnit,
			&ingredientsJSON,
		); err != nil {
			return nil, fmt.Errorf("failed to scan user routine product row: %w", err)
		}

		if err := json.Unmarshal(ingredientsJSON, &product.IngredientsWithImpact); err != nil {
			return nil, fmt.Errorf("failed to unmarshal ingredients JSON: %w", err)
		}

		product.Source = "internal"

		if rp, exists := routineMap[routineType]; exists {
			rp.Products = append(rp.Products, &product)
		} else {
			routineMap[routineType] = &domain.UserRoutineProduct{
				Products:    []*domain.Product{&product},
				RoutineType: routineType,
				CreatedAt:   routineCreatedAt,
			}
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating user routine product rows: %w", err)
	}

	var result []*domain.UserRoutineProduct
	for _, rp := range routineMap {
		result = append(result, rp)
	}

	return result, nil
}
