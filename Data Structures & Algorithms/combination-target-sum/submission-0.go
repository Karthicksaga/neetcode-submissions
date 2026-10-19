func backtrackCombination(nums []int, remaining int, start int, path []int, result *[][]int){

	if remaining == 0 {
		*result = append(*result, append([]int{}, path...))
		return 
	}

	if remaining < 0{
		return 
	}


	for i := start; i < len(nums); i += 1 {

		path = append(path, nums[i])
		backtrackCombination(nums, remaining - nums[i], i, path, result)
		path = path[:len(path) - 1]
	}

}

func combinationSum(nums []int, target int) [][]int {

 	result := [][]int{}
    backtrackCombination(nums, target, 0, []int{}, &result)
    return result
    
}
