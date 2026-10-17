class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        
        numset = set(nums)
        result = 0
        for n in nums:
           
            if (n - 1) not in numset:
                longest = 1
                while n + longest in nums:
                    longest += 1
            
                result = max(longest, result)
        return result
            