from typing import List
from collections import deque


def twoColorable(edges: List[List[int]]) -> bool:
    colors = [False for _ in range(len(edges))]
    colors[0] = True
    stack = deque()
    stack.append(0)

    while len(stack) > 0:
        node = stack.pop()
        for neighbor in edges[node]:
            if colors[neighbor] is False:
                colors[neighbor] = not colors[node]
                stack.append(neighbor)
            elif colors[neighbor] == colors[node]:
                return False
    return True
