# This is an input class. Do not edit.
class BinaryTree:
    def __init__(self, value, left=None, right=None):
        self.value = value
        self.left = left
        self.right = right


def findNodesDistanceK(tree, target, k):
    nodes_distance_k = []
    find_nodes_distance_k(tree, target, k, nodes_distance_k)
    return nodes_distance_k


def find_nodes_distance_k(tree, target, k, nodes_distance_k):
    if tree is None:
        return -1
    if tree.value == target:
        add_subtree_nodes_at_distance_k(tree, 0, k, nodes_distance_k)
        return 1
    left_distance = find_nodes_distance_k(tree.left, target, k, nodes_distance_k)
    right_distance = find_nodes_distance_k(tree.right, target, k, nodes_distance_k)

    if left_distance == k or right_distance == k:
        nodes_distance_k.append(tree.value)

    if left_distance != -1:
        add_subtree_nodes_at_distance_k(
            tree.right, left_distance + 1, k, nodes_distance_k
        )
        return left_distance + 1
    if right_distance != -1:
        add_subtree_nodes_at_distance_k(
            tree.left, right_distance + 1, k, nodes_distance_k
        )
        return right_distance + 1
    return -1


def add_subtree_nodes_at_distance_k(tree, distance, k, nodes_distance_k):
    if tree is None:
        return
    if distance == k:
        nodes_distance_k.append(tree.value)
        return
    add_subtree_nodes_at_distance_k(tree.left, distance + 1, k, nodes_distance_k)
    add_subtree_nodes_at_distance_k(tree.right, distance + 1, k, nodes_distance_k)
