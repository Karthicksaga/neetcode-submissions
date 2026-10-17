class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        
        if len(s) != len(t):
            return False
            
        str1dict, str2dict = dict(), dict()

        for char in s:
            str1dict[char] = 1 +  str1dict.get(char, 0)
        
        for char in t:
            str2dict[char] = 1 +  str2dict.get(char, 0)
        
        for key, value in str1dict.items():
            if key not in str2dict or value != str2dict[key]:
                return False
        return True
                