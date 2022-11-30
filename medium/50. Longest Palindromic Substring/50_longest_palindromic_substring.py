def longestPalindromicSubstring(string: str) -> str:
    longest_indexes = [0, 1]

    for i in range(1, len(string)):
        odd = get_current_longest_palindromix_indexes(string, i - 1, i + 1)
        even = get_current_longest_palindromix_indexes(string, i - 1, i)
        longest = max(odd, even, key=lambda x: x[1] - x[0])
        longest_indexes = max(longest, longest_indexes, key=lambda x: x[1] - x[0])
    return string[longest_indexes[0] : longest_indexes[1]]


def get_current_longest_palindromix_indexes(string, left, right):
    while left >= 0 and right < len(string):
        if string[left] != string[right]:
            break
        left -= 1
        right += 1
    return [left + 1, right]
