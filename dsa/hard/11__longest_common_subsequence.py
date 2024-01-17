from typing import List


def longestCommonSubsequence(s1, s2) -> List[str]:
    lengths = [[0 for x in range(len(s1) + 1)] for y in range(len(s2) + 1)]

    for i in range(1, len(s2) + 1):
        for j in range(1, len(s1) + 1):
            if s2[i - 1] == s1[j - 1]:
                lengths[i][j] = lengths[i - 1][j - 1] + 1
            else:
                lengths[i][j] = max(lengths[i - 1][j], lengths[i][j - 1])
    return make_longest_common_subsequence(lengths, s1)


def make_longest_common_subsequence(lengths, s):
    sequence = []
    i = len(lengths) - 1
    j = len(lengths[0]) - 1
    while i != 0 and j != 0:
        if lengths[i][j] == lengths[i - 1][j]:
            i -= 1
            continue
        if lengths[i][j] == lengths[i][j - 1]:
            j -= 1
            continue
        sequence.append(s[j - 1])
        i -= 1
        j -= 1

    return list(reversed(sequence))
