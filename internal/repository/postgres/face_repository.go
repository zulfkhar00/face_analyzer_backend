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
