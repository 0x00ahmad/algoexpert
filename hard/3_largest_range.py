from typing import List


def largestRange(array: List[int]) -> List[int]:
    """
    Time: O(n)
    Space: O(n)
    """
    hash_table = {}
    for num in array:
        hash_table[num] = True
    longest_range = []
    longest_length = 0
    for num in array:
        if not hash_table[num]:
            continue
        hash_table[num] = False
        current_length = 1
        left = num - 1
        right = num + 1
        while left in hash_table:
            hash_table[left] = False
            current_length += 1
            left -= 1
        while right in hash_table:
            hash_table[right] = False
            current_length += 1
            right += 1
        if current_length > longest_length:
            longest_length = current_length
            longest_range = [left + 1, right - 1]
    return longest_range
