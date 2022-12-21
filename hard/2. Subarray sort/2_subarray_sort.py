from typing import List


def subarraySort(array: List[int]) -> List[int]:
    """
    time: O(n)
    space: O(1)
    """

    # find the first number that is out of order from the beginning
    minOutOfOrder = float("inf")
    for i in range(len(array) - 1):
        if array[i] > array[i + 1]:
            minOutOfOrder = min(minOutOfOrder, array[i + 1])

    # if the array is already sorted, return [-1, -1]
    if minOutOfOrder == float("inf"):
        return [-1, -1]

    # find the first number that is out of order from the end
    maxOutOfOrder = float("-inf")
    for i in reversed(range(1, len(array))):
        if array[i] < array[i - 1]:
            maxOutOfOrder = max(maxOutOfOrder, array[i - 1])

    # find the correct position of the minOutOfOrder and maxOutOfOrder
    subarrayLeftIdx = 0
    while minOutOfOrder >= array[subarrayLeftIdx]:
        subarrayLeftIdx += 1

    subarrayRightIdx = len(array) - 1
    while maxOutOfOrder <= array[subarrayRightIdx]:
        subarrayRightIdx -= 1

    return [subarrayLeftIdx, subarrayRightIdx]


def test_1():
    assert subarraySort([1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19]) == [3, 9]
