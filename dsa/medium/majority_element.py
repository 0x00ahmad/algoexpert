from collections import Counter


def majorityElement(array):
    return max(Counter(array).items(), key=lambda x: x[1])[0]
