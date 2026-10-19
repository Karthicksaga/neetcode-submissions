func longestCommonPrefix(strs []string) string {

	prefix := strs[0]

	for i := 1; i < len(strs); i += 1{

		for !strings.HasPrefix(strs[i], prefix){
			prefix = prefix[:len(prefix) - 1]
			fmt.Println("Trim Values: ", strs[i])
			if prefix == ""{
				return prefix
			}
		}
	}

	return prefix
    
}
