func reverse(nums []int, start, end int) []int{

	for start < end{
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}

	return nums
}


func rotate(nums []int, k int) {

	if len(nums) == 0{
		return
	}

	n := len(nums)

	k = k % n
	
	nums = reverse(nums, 0, n-k - 1)
	nums = reverse(nums, (n-k), n - 1)
	nums = reverse(nums, 0, n-1)

	fmt.Println("Rotated Array: ", nums)

}
