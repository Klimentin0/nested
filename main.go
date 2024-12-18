package main

func getNameCounts(names []string) map[rune]map[string]int {
	nested := make(map[rune]map[string]int)
	//  sumName := make(map[string]int)
	for _, user := range names {
		firstRune := []rune(user)[0]
		if nested[firstRune] == nil {
			nested[firstRune] = make(map[string]int)
		}

		nested[firstRune][user]++
	}
	return nested
}
