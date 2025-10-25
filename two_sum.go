func twoSum(nums []int, target int) []int {
	comp := make(map[int]int)

	for i, v := range nums {
		if _, exists := comp[v]; exists {
			return []int{comp[v], i}
		}
		comp[target-v] = i
	}
	return []int{-1, -1}
}
