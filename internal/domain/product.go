package domain

import "time"

// Product represents a skincare product
type Product struct {
	ID                    string             `json:"id"`                  // Corresponds to p.id AS product_id (UUID from DB)
	Barcode               string             `json:"barcode"`             // Corresponds to p.code AS barcode (string from DB)
	ProductName           string             `json:"productName"`         // Corresponds to p.product_name
	Brand                 string             `json:"brand"`               // Corresponds to p.brands AS brand
	ProductQuantity       float64            `json:"productQuantity"`     // Corresponds to p.product_quantity
	ProductQuantityUnit   string             `json:"productQuantityUnit"` // Corresponds to p.product_quantity_unit
	Ingredients           []string           `json:"ingredients"`         // Corresponds to json_agg(DISTINCT i.ingredient_name) AS ingredients
	IngredientsWithImpact []IngredientImpact `json:"ingredientsWithImpacts"`
	Source                string             `json:"source"` // This field is application-specific, not from DB, so will be set by service/handler
}

type IngredientImpact struct {
	IngredientName               string  `json:"ingredient_name"`
	BlackheadsImpact             string  `json:"blackheads_impact"`
	BlackheadsConfidence         float64 `json:"blackheads_confidence"`
	DarkCirclesOnFaceImpact      string  `json:"dark_circles_on_face_impact"`
	DarkCirclesOnFaceConfidence  float64 `json:"dark_circles_on_face_confidence"`
	DarkSpotsOnFaceImpact        string  `json:"dark_spots_on_face_impact"`
	DarkSpotsOnFaceConfidence    float64 `json:"dark_spots_on_face_confidence"`
	DrySkinImpact                string  `json:"dry_skin_impact"`
	DrySkinConfidence            float64 `json:"dry_skin_confidence"`
	DullSkinImpact               string  `json:"dull_skin_impact"`
	DullSkinConfidence           float64 `json:"dull_skin_confidence"`
	EyeBagsImpact                string  `json:"eye_bags_impact"`
	EyeBagsConfidence            float64 `json:"eye_bags_confidence"`
	FaceRednessImpact            string  `json:"face_redness_impact"`
	FaceRednessConfidence        float64 `json:"face_redness_confidence"`
	ForeheadWrinklesImpact       string  `json:"forehead_wrinkles_impact"`
	ForeheadWrinklesConfidence   float64 `json:"forehead_wrinkles_confidence"`
	HormonalAcneImpact           string  `json:"hormonal_acne_impact"`
	HormonalAcneConfidence       float64 `json:"hormonal_acne_confidence"`
	LargePoresOnFaceImpact       string  `json:"large_pores_on_face_impact"`
	LargePoresOnFaceConfidence   float64 `json:"large_pores_on_face_confidence"`
	OilySkinImpact               string  `json:"oily_skin_impact"`
	OilySkinConfidence           float64 `json:"oily_skin_confidence"`
	RazorBumpsImpact             string  `json:"razor_bumps_impact"`
	RazorBumpsConfidence         float64 `json:"razor_bumps_confidence"`
	RoughTextureOnFaceImpact     string  `json:"rough_texture_on_face_impact"`
	RoughTextureOnFaceConfidence float64 `json:"rough_texture_on_face_confidence"`
	SebaceousFilamentsImpact     string  `json:"sebaceous_filaments_impact"`
	SebaceousFilamentsConfidence float64 `json:"sebaceous_filaments_confidence"`
	UnderEyeWrinklesImpact       string  `json:"under_eye_wrinkles_impact"`
	UnderEyeWrinklesConfidence   float64 `json:"under_eye_wrinkles_confidence"`
}

type UserRoutineProduct struct {
	Products    []*Product
	RoutineType string    // from user_skincare_routine.routine_type, e.g.: morning, evening, both
	CreatedAt   time.Time // from user_skincare_routine.created_at
}
