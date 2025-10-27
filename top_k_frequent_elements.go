func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}
	buckets := make([][]int, len(nums)+1)
	for v, f := range freq {
		buckets[f] = append(buckets[f], v)
	}

	answer := []int{}
	for i := len(buckets) - 1; i > 0; i-- {
		for j := 0; j < len(buckets[i]); j++ {
			answer = append(answer, buckets[i][j])
			if len(answer) == k {
				return answer
			}
		}
	}
	return answer
}
