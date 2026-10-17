func topKFrequent(nums []int, k int) []int {

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
result := []int{}
for i:= len(buckets) - 1; i > 0; i--{

	for _, value := range buckets[i]{

		result = append(result, value)

		if len(result) == k{
			return result 
		}
	}
		
}

return result
}
