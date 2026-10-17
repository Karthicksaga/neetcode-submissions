class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        
        index ={}

        for idx, ele in enumerate(numbers):
            index[ele] = idx
        
        for idx, element in enumerate(numbers):

            curSum = target-element

            if curSum in index:
                return[ idx + 1, index[curSum] + 1]
        
        return []