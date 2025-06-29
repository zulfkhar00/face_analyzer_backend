package domain

// Face represents a user face
type Face struct {
	UID              string
	Probabilities    map[string]float64
	OverallScore     float32
	OverallCondition string
}
