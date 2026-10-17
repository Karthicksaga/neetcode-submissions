func topKFrequent(nums []int, k int) []int {

if len(nums) == k{
	return nums
}

//compute the frequency 
frequency := make(map[int]int)

for _, num := range nums{
	frequency[num]++
}

// prepare the bucket sort 
buckets :=  make([][]int, len(nums) + 1)

for num, count := range frequency{
	buckets[count] = append(buckets[count], num)
}

//take the elements 
result := make([]int, 0, k)
for i:= len(buckets) - 1; i > 0 && k > 0; i--{

	for _, value := range buckets[i]{

		result = append(result, value)

		if len(result) == k{
			return result 
		}
	}
		
}

return result
}
