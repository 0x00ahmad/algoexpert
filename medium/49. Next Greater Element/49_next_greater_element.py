from typing import List


def nextGreaterElement(array: List[int]) -> List[int]:
    array_len = len(array)
    result = [-1] * array_len
    stack = []
    for index in range(2 * array_len - 1, -1, -1):
        modded_index = index % array_len
        while len(stack) > 0:
            if stack[-1] <= array[modded_index]:
                stack.pop()
                continue
            result[modded_index] = stack[len(stack) - 1]
            break
        stack.append(array[modded_index])
    return result
