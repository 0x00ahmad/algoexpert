from typing import List


def sortStack(stack: List[int]) -> List[int]:
    if not len(stack):
        return stack

    top = stack.pop()

    sortStack(stack)

    insert_in_sorted_order(stack, top)

    return stack


def insert_in_sorted_order(stack: List[int], value: int):
    if not len(stack) or stack[-1] <= value:
        stack.append(value)
        return

    top = stack.pop()

    insert_in_sorted_order(stack, value)

    stack.append(top)
