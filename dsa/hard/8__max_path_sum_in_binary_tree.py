def maxPathSum(tree):
    _, maxSum = find_max_sum(tree)
    return maxSum


def find_max_sum(tree):
    if tree is None:
        return (0, float("-inf"))

    left_max_sum_as_branch, left_max_path_sum = find_max_sum(tree.left)
    right_max_sum_as_branch, right_max_path_sum = find_max_sum(tree.right)
    max_child_sum_as_branch = max(left_max_sum_as_branch, right_max_sum_as_branch)

    current_value = tree.value
    max_sum_as_branch = max(current_value, current_value + max_child_sum_as_branch)
    max_sum_as_root_node = max(
        max_sum_as_branch,
        left_max_sum_as_branch + current_value + right_max_sum_as_branch,
    )
    max_path_sum = max(left_max_path_sum, right_max_path_sum, max_sum_as_root_node)

    return (max_sum_as_branch, max_path_sum)
