func groupAnagrams(strs []string) [][]string {
	inter := make(map[[26]int][]string)

	for _, s := range strs {
		var count [26]int
		for _, c := range s {
			count[c-'a']++
		}
		inter[count] = append(inter[count], s)
	}
	
	var answer [][]string
	for _, g := range inter {
		answer = append(answer, g)
	}
	return answer
}
