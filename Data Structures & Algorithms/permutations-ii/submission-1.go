
func backtrackPermutation(nums []int, used []bool, path []int,result *[][]int){

	if len(nums) == len(path){
		*result = append(*result, append([]int{}, path...))
		return  
	}


	for i := 0; i < len(nums); i += 1{

		if used[i]{
			continue
		}
		if i > 0 && nums[i] == nums[i - 1] && !used[i - 1]{
			continue
		}

		used[i] = true
		path = append(path, nums[i])
		backtrackPermutation(nums, used, path, result)
		path = path[:len(path) - 1]
		used[i] = false
	}
}

func permuteUnique(nums []int) [][]int {

	sort.Ints(nums)
	result := [][]int{}
    used := make([]bool, len(nums))

    backtrackPermutation(nums, used, []int{}, &result)
    return result


}
