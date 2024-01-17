from typing import List


def maxSumIncreasingSubsequence(array):
    sequences: List[int | None] = [None for x in array]
    sums = array[:]
    max_sum_idx = 0
    for i in range(len(array)):
        current_num = array[i]
        for j in range(0, i):
            other_num = array[j]
            if other_num < current_num and sums[j] + current_num >= sums[i]:
                sums[i] = sums[j] + current_num
                sequences[i] = j
        if sums[i] >= sums[max_sum_idx]:
            max_sum_idx = i
    return [sums[max_sum_idx], make_sequence(array, sequences, max_sum_idx)]


def make_sequence(array, sequences, max_sum_idx):
    sequence = []
    while max_sum_idx is not None:
        sequence.append(array[max_sum_idx])
        max_sum_idx = sequences[max_sum_idx]
    return list(reversed(sequence))
