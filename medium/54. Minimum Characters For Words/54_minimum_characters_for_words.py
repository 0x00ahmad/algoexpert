from collections import Counter
from typing import List


def minimumCharactersForWords(words: List[str]) -> List[str]:
    """
    1. make a dict with {char: count} where count is the max count btw all
    of the words' characters passed
    2. return the final list from that dict f (frequency) many times over each character
    """
    word_count = {}
    for word in words:
        current_word_count = Counter(word)
        for (c, f) in current_word_count.items():
            word_count[c] = f if not c in word_count else max(word_count[c], f)
    return [
        character for (character, count) in word_count.items() for _ in range(count)
    ]
