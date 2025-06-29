package scripts

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

func RunFaceAnalysisPython(imagePath string) (string, map[string]float64, error) {
	cmd := exec.Command("python", "./scripts/face_analysis.py", imagePath)

	out, err := cmd.Output()
	if err != nil {
		return "", nil, err
	}
	outStr := string(out)
	jsonStart := strings.Index(outStr, "{")
	jsonEnd := strings.LastIndex(outStr, "}") + 1

	if jsonStart == -1 || jsonEnd == 0 || jsonEnd <= jsonStart {
		return "", nil, fmt.Errorf("couldn't find valid JSON in output: %s", outStr)
	}

	jsonStr := outStr[jsonStart:jsonEnd]

	var result struct {
		PredictedSkin string             `json:"predicted_skin"`
		Probabilities map[string]float64 `json:"probabilities"`
	}
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return "", nil, err
	}

	return result.PredictedSkin, result.Probabilities, nil
}
