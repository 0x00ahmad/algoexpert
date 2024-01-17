# This is an input class. Do not edit.
class BinaryTree:
    def __init__(self, value, left=None, right=None):
        self.value = value
        self.left = left
        self.right = right


def is_symmetrical(left_subtree, right_subtree):
    if left_subtree is None and right_subtree is None:
        return True
    if left_subtree is None or right_subtree is None:
        return False
    if left_subtree.value != right_subtree.value:
        return False
    is_left_symmetrical = is_symmetrical(left_subtree.left, right_subtree.right)
    is_right_symmetrical = is_symmetrical(left_subtree.right, right_subtree.left)
    return is_left_symmetrical and is_right_symmetrical


def symmetricalTree(tree):
    if tree is None:
        return True
    return is_symmetrical(tree.left, tree.right)
