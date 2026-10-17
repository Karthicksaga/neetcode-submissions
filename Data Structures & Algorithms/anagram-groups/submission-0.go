func groupAnagrams(strs []string) [][]string {

	seen := map[[26]int][]string{}

	for _, str := range strs{

		var char [26]int
		for _, s := range str{
			char[s - 'a'] += 1
		}

		seen[char] = append(seen[char], str)

	}

	result := [][]string{}
	for _, anagrams := range seen{
		result = append(result,anagrams)
	}

	return result

}
