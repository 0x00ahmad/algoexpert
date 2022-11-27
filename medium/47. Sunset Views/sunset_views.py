from typing import List, Literal
from collections import deque

DIRECTION = Literal["EAST", "WEST"]
BUILDINGS = List[int]


def sunsetViews(buildings: BUILDINGS, direction: DIRECTION) -> BUILDINGS:
    valid_views = deque()
    lb = len(buildings)
    current_peak = -float("inf")
    if direction == "EAST":
        for i in range(lb - 1, -1, -1):
            building = buildings[i]
            if building > current_peak:
                # we are only appending *if* the building coming after
                valid_views.appendleft(i)
                current_peak = building
    elif direction == "WEST":
        for i in range(lb):
            building = buildings[i]
            if building > current_peak:
                # we are only appending *if* the building coming after
                valid_views.append(i)
                current_peak = building
    else:
        raise ValueError("Only 'EAST' and 'WEST' are allowed as directions")
    return list(valid_views)
