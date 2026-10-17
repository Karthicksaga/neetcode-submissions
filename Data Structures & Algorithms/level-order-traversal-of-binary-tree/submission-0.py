# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from collections import deque

class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:

        if not root:
            return []

        result = []
        queue = deque([root])
        while queue:
            levelLength = len(queue)
            level = []
            while levelLength > 0:
                
                currentNode =  queue.popleft()
                level.append(currentNode.val)

                if currentNode.left:
                    queue.append(currentNode.left)
                if currentNode.right:
                    queue.append(currentNode.right)
                
                levelLength -= 1
                
            result.append(level)
            print(result)
        
        return result

        

        