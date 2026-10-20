func moveZeroes(nums []int) {
	
	pos := 0 

	for i := 0; i < len(nums); i += 1{
		if nums[i] != 0{
			nums[pos] = nums[i]
			pos += 1
		}
	}

	for pos < len(nums){
		nums[pos] = 0
		pos += 1
	}
	
	
}
