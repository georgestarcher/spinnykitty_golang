package utilities

// Utility function to calc Mean value
func CalcMean(data []int) float64 {

	if len(data) == 0 {
		return 0
	}

	var sum float64
	for _, d := range data {
		sum += float64(d)
	}
	return sum / float64(len(data))
}
