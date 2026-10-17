func subarraySum(nums []int, k int) int {

	result, currentSum := 0, 0 
	prefixSum := map[int]int{0: 1}


	for _, num := range nums{

		currentSum += num
		difference := currentSum - k 

		result += prefixSum[difference]

		prefixSum[currentSum] += 1 

	}

	return result
}
