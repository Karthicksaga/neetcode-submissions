func maxArea(heights []int) int {

	l, r := 0, len(heights) - 1 
	maxArea := 0

	for l < r{

		heightMin := min(heights[l], heights[r])
		area := (r - l) * heightMin
		maxArea = max(maxArea,area)

		if heights[l] > heights[r]{
			r -= 1
		}else{
			l += 1
		}

	}

	return maxArea

}
