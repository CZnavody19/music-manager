package utils

import "math"

func NumberSimilarity(a, b float64) float64 {
	den := math.Max(math.Abs(a), math.Abs(b))
	if den == 0 {
		return 1.0
	}
	q := math.Abs((a - b) / den)
	return 1.0 - math.Min(q, 1.0)
}

func IntPtr(n int) *int {
	return &n
}
