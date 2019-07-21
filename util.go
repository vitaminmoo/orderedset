package orderedset

// removeIndex nukes an index from a slice, and returns it as a new slice
func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
