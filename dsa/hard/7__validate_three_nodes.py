"""
constraints:

1. One of n1 or n3 is an ancestor of n2
2. The other of the two nodes is a descendant of n2,
    example if n1 is ancestor of n2, then we have to check if n3 is a descendant of nodeTwo
        if n3 is a descendant of n2, then we have to check if n1 is an ancestor of nodeTwo

A descendant is defined as a node that is on the path from the given node down to a leaf node
an ancestor is defined as a node that is on the path from the given node up to the root node

There is one gurantee, which is that the given nodes are all unique and there will not be any null values

"""


# This is an input class. Do not edit.
class BST:
    def __init__(self, value, left=None, right=None):
        self.value = value
        self.left = left
        self.right = right


# time O(d) | space O(1) | iterative
def validateThreeNodes(n1: BST, n2: BST, n3: BST) -> bool:
    is_node_two_found = False
    node = n1

    while node is not None:
        if node == n2:
            is_node_two_found = True
        if node.value < n3.value:
            node = node.right
        elif node.value > n3.value:
            node = node.left
        else:
            if is_node_two_found:
                return True
            break

    is_node_two_found = False
    node = n3
    while node is not None:
        if node == n2:
            is_node_two_found = True
        if node.value < n1.value:
            node = node.right
        elif node.value > n1.value:
            node = node.left
        else:
            return is_node_two_found

    return False
