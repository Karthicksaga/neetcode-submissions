class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:


        charSet = set()
        l, r   = 0, 0 
        maxLength = 0
        while r < len(s):

            while s[r] in charSet:
                charSet.remove(s[l])
                l = l + 1
            
            charSet.add(s[r])

            maxLength = max(maxLength, r - l + 1)

            r = r + 1
        return maxLength