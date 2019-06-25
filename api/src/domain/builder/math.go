package builder

import "math"

// Round Arredonta o número com dois decimais
func Round(f float64) float64 {
	return math.Round(f*100) / 100
}
