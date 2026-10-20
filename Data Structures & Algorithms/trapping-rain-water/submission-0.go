func trap(heights []int) int {

	if len(heights) == 0{
		return 0
	}

	l, r := 0, len(heights) - 1 
	leftMax, rightMax := 0 , 0
	water := 0

	for l < r{

		if heights[l] <= heights[r]{

			if heights[l] > leftMax{
				leftMax = heights[l]
			}else{
				water += leftMax - heights[l]
			}
			l += 1

		} else{

			if heights[r] > rightMax{
				rightMax = heights[r]
			}else{
				water += rightMax - heights[r]
			}

			r -= 1
		}
	}

	return water
}
