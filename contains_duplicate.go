func hasDuplicate(nums []int) bool {
	seen := make(map[int]struct{})

	for _, v := range nums {
		if _, exists := seen[v]; exists {
			return true
		}
		seen[v] = struct{}{}
	}
	return false
}
