from typing import List


def bestSeat(seats: List[int]):
    best_seat = -1
    max_space = 0
    i = 0

    while i < len(seats):
        j = i + 1
        while j < len(seats) and seats[j] == 0:
            j += 1
            available_space = j - i - 1
            if available_space > max_space:
                best_seat = (j + i) // 2
                max_space = available_space
        i = j
    return best_seat
