from typing import List


def is_out_of_bounds(row: int, col: int, height: int, width: int) -> bool:
    return row < 0 or row > height or col < 0 or col > width


def zigzagTraverse(array: List[List[int]]) -> List[int]:
    out = []
    height = len(array) - 1
    width = len(array[0]) - 1
    row, col = 0, 0
    going_down = True

    while not is_out_of_bounds(row, col, height, width):
        out.append(array[row][col])
        if going_down:
            if col == 0 or row == height:
                going_down = False
                if row == height:
                    col += 1
                else:
                    row += 1
            else:
                row += 1
                col -= 1
            continue
        if row == 0 or col == width:
            going_down = True
            if col == width:
                row += 1
            else:
                col += 1
        else:
            row -= 1
            col += 1
    return out
