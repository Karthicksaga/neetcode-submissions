func maxProduct(nums []int) int {
    
	prefix := 1
	suffix := 1 
	result := nums[0]

	n := len(nums)
	for i := 0; i < n; i += 1{
		if prefix == 0{
			prefix = 1
		}

		if suffix == 0{
			suffix = 1
		}
		
		prefix *= nums[i]
		suffix *= nums[n-1-i]

		result = max(result, max(prefix, suffix))


	}

	return result
}
