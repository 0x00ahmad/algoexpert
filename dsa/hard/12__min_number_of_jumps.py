from typing import List


# time O(n) | space O(1)
def minNumberOfJumps(array: List[int]):
    if len(array) == 1:
        return 0
    jumps = 0
    max_reach = array[0]
    steps = array[0]
    for i in range(1, len(array) - 1):
        max_reach = max(max_reach, array[i])
        steps -= 1
        if steps == 0:
            jumps += 1
            steps = max_reach
    return jumps + 1
