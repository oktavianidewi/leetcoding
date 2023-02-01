class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def binaryTreePaths(self, root: Optional[TreeNode]) -> List[str]:
        return 9


if __name__ == "__main__" :
    some_dict={"user":"macro","address":"some street","age":28}
    tree = TreeNode(val=1)
    tree.left = TreeNode(left=2)
    tree.left.right = TreeNode(right=5)
    tree.right = TreeNode(right=3)
    print(tree)
