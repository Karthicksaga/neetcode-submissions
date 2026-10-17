func twoSum(nums []int, target int) []int {

    seen := map[int]int{}

    for index, num := range nums{

        complement := target - num
        if val, ok := seen[complement]; ok{
            return []int{val, index}
        }

        seen[num] = index
    }
    
    return []int{}
}
