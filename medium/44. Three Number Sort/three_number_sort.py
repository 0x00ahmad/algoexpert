from typing import List, Tuple


ORDER = Tuple[float, float, float]


def threeNumberSort(array: List, order: ORDER):
    """
    The DNF Algorithm
    """
    assert len(order) == 3
    la = len(array)
    low = 0
    mid = 0
    high = la - 1

    while mid <= high:
        if array[mid] == order[0]:
            array[low], array[mid] = array[mid], array[low]
            low += 1
            mid += 1
        elif array[mid] == order[1]:
            mid += 1
        elif array[mid] == order[2]:
            array[mid], array[high] = array[high], array[mid]
            high -= 1

    return array
