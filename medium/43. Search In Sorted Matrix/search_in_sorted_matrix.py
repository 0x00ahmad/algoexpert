def searchInSortedMatrix(matrix, target: int):
    out: list[int] = [-1, -1]
    for row_idx in range(len(matrix)):
        if matrix[row_idx][-1] < target:
            continue  # we have to find that row whom's last element is larger than our target
        for ri in [row_idx - 1, row_idx]:
            try:
                out[1] = matrix[ri].index(target)
                out[0] = ri
            except ValueError:
                continue
    return out
