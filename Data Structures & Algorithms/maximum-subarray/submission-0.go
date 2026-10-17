func maxSubArray(nums []int) int {

	currentSum, maxSum := 0, nums[0] 

	for _, num := range nums{

		currentSum = currentSum + num

		if num > currentSum{
			currentSum = num
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
    
	return maxSum
}
