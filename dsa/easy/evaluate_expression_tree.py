import math

# This is an input class. Do not edit.


class BinaryTree:
    def __init__(self, value, left=None, right=None):
        self.value = value
        self.left = left
        self.right = right


ADD = -1
SUBTRACT = -2
DIVISION = -3
MULTIPLICATION = -4


def evaluate(sub_tree: BinaryTree) -> int:
    if sub_tree.left is None and sub_tree.right is None:
        return sub_tree.value  # leaf node
    if sub_tree.value == ADD:
        return evaluate(sub_tree.left) + evaluate(sub_tree.right)
    elif sub_tree.value == SUBTRACT:
        return evaluate(sub_tree.left) - evaluate(sub_tree.right)
    elif sub_tree.value == DIVISION:
        return evaluate(sub_tree.left) / evaluate(sub_tree.right)
    return evaluate(sub_tree.left) * evaluate(sub_tree.right)


def evaluateExpressionTree(tree):
    return evaluate(tree)


def main():
    tree = BinaryTree(
        -1,
        BinaryTree(
            -2, BinaryTree(-4, BinaryTree(2), BinaryTree(3)), BinaryTree(2)
        ),
        BinaryTree(-3, BinaryTree(8), BinaryTree(3)),
    )
    assert evaluateExpressionTree(tree) == 6

    tree = BinaryTree(
        -3,
        BinaryTree(9),
        BinaryTree(-2, BinaryTree(4), BinaryTree(6)),
    )
    assert evaluateExpressionTree(tree) == 4.5


if __name__ == "__main__":
    main()
