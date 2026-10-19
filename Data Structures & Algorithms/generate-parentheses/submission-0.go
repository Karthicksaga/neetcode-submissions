func backtrack(n int, open, close int, path string, result *[]string){

	if len(path) == 2 * n{
		*result = append(*result, path)
		return 
	}


	if open < n{
		backtrack(n, open + 1, close, path + "(", result )
	}

	if close < open{
		backtrack(n, open, close + 1, path + ")", result )
	}


}

func generateParenthesis(n int) []string {

	result := []string{}
	backtrack(n, 0, 0, "", &result)

	return result


}
