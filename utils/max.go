package utils

func max(a, b int) int {
// Remove the unused function to fix the warning
	if a < b {
		return b
	}
	return a
}
