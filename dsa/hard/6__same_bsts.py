from typing import List


def parse_array_to_bst_map(array: List[int]) -> dict:
    bst = {}
    for num in array:
        j = 0
        while j in bst:
            if num >= bst[j]:
                j = 2 * j + 1
            else:
                j = 2 * j + 2
        bst[j] = num
    return bst


def sameBsts(a1: List[int], a2: List[int]) -> bool:
    """
    as the only constraint is that we can't form a bst, it didn't stop us from
    parsing a false map of a bst. So we can just compare the maps representing
    {index: value} of the bst.
    """
    if len(a1) != len(a2):
        return False
    if len(a1) == 0:
        return True
    if a1[0] != a2[0]:
        return False

    bst1 = parse_array_to_bst_map(a1)
    bst2 = parse_array_to_bst_map(a2)

    return bst1 == bst2
