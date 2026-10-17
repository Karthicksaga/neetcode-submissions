# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from collections import deque
class Solution:
    def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
        
        result = []
        if not root:
            return []
        
        queue = deque([root])
        while queue:
            length = len(queue)

            while length > 0:
                node = queue.popleft()
                if length == 1:
                    print(node.val)
                    result.append(node.val)
                
                if node.left:
                    queue.append(node.left)
                if node.right:  
                    queue.append(node.right)

                length -=1
            print(result) 
        return result 
