package utils

// TotalPrice calculates the total price
func TotalPrice(q float64, p float64) float64 {
	total := float64(q) * float64(p)

	return total
}
