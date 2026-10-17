class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        


        result = 0
        countS = {}
        l = 0
        maxF = 0

        for r in range(len(s)):

            countS[s[r]] = 1 + countS.get(s[r], 0)
            maxF = max(maxF, countS[s[r]])


            while (r - l + 1) - maxF > k:
                countS[s[l]] -= 1
                l += 1

            result = max(result, r - l + 1)

        return result