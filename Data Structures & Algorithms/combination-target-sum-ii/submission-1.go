func backtrackCombination(nums []int, remaining int, start int, path []int, result *[][]int){


	if remaining == 0 {

		*result = append(*result, append([]int{}, path...))
		return 
	}

	if remaining < 0{
		return 
	}

	for i := start ; i < len(nums); i += 1{

		if i > start && nums[i] == nums[i - 1]{
			continue
		}
		

		path = append(path, nums[i])
		backtrackCombination(nums, remaining - nums[i], i + 1,  path, result)
		path = path[:len(path) - 1]
	}
} 

func combinationSum2(candidates []int, target int) [][]int {
 sort.Ints(candidates)
    result := [][]int{}
    backtrackCombination(candidates, target, 0, []int{}, &result)
    return result
}
