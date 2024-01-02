package uniqueUtil

func UniqueInt64Slice(slice []int64) []int64 {
	seen := make(map[int64]bool)
	var result []int64

	for _, v := range slice {
		if _, ok := seen[v]; !ok {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}
