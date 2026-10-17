class Solution:
    def isValid(self, s: str) -> bool:

        stack = []
        for c in s:
            if c in ")}]" and len(stack) > 0:
                popValue = stack.pop()
                print("PopValue", popValue)
                if (popValue != "[" and c == "]") or (popValue != "{" and c == 
                "}") or (popValue != "(" and c==")"):
                    return False 
            else:
                stack.append(c)
        
        return True if not stack else False

            