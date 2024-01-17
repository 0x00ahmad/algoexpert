def transposeMatrix(matrix):
    out = [[0 for _ in range(len(matrix))] for _ in range(len(matrix[0]))]
    for i in range(len(matrix)):
        for j in range(len(matrix[0])):
            out[j][i] = matrix[i][j]
    return out


def test():
    matrix = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
    assert transposeMatrix(matrix) == [[1, 4, 7], [2, 5, 8], [3, 6, 9]]
    print("Test passed")


if __name__ == "__main__":
    test()
