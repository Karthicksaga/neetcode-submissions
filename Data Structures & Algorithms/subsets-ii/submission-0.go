func backtrack(nums []int, start int, path []int, result *[][]int){

	*result = append(*result, append([]int{}, path...))

	for i := start; i < len(nums); i += 1{

		if i > start && nums[i] == nums[i - 1]{
			continue
		}
	
		path = append(path, nums[i])
		backtrack(nums, i + 1, path, result)
		path = path[:len(path) - 1]
	}
}



func subsetsWithDup(nums []int) [][]int {

	sort.Ints(nums)
result := [][]int{}
	backtrack(nums, 0, []int{}, &result)
	return result
}

