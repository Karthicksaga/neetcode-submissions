class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
         
        duplicateSet = set()

        for num in nums:
            if num in duplicateSet:
                return True 

            duplicateSet.add(num)
        return False