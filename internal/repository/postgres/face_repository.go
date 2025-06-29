package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"

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
