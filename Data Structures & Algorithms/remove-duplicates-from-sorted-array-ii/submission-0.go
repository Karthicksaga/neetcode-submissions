func removeDuplicates(nums []int) int {

	l := 0 
	for r := 0; r < len(nums); r += 1{

		if l < 2 || nums[r] != nums[l - 2]{
			nums[l] = nums[r]
			l += 1
		}
	}
    
	return l
}