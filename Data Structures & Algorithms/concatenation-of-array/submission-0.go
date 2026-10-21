func getConcatenation(nums []int) []int {
	n := len(nums)
	answer := make([]int, 2 * n) 

	for i, num :=  range nums{
		answer[i], answer[i + n ] = num, num
	}

	return answer
}
