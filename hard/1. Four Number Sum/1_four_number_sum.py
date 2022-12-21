def fourNumberSum(array, targetSum):
    """
    time: O(n^2)
    space: O(n^2)
    """
    out = []

    # create a dictionary to store the sum of two numbers
    # and the pair of numbers that add up to that sum
    sums = {}

    for i in range(1, len(array) - 1):
        for j in range(i + 1, len(array)):
            currentSum = array[i] + array[j]
            difference = targetSum - currentSum

            if difference in sums:
                for pair in sums[difference]:
                    out.append(pair + [array[i], array[j]])

        for k in range(0, i):
            currentSum = array[i] + array[k]
            if currentSum not in sums:
                sums[currentSum] = [[array[k], array[i]]]
            else:
                sums[currentSum].append([array[k], array[i]])

    return out


def test_1():
    assert fourNumberSum([7, 6, 4, -1, 1, 2], 16) == [[7, 6, 4, -1], [7, 6, 1, 2]]
