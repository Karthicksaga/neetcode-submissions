class Solution:
    def maxArea(self, heights: List[int]) -> int:
        
        l, r = 0, len(heights) - 1

        result = 0
        area = 0
        while l < r:
            area = (r - l) * min(heights[l], heights[r])
            result = max(result, area)
            if heights[r] > heights[l]:
                l += 1
            elif heights[r] <= heights[l]:
                r -= 1

        return result
            
                    