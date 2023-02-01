from typing import List


def semordnilap(words: List[str]) -> List:
    """
    Given a list of words, find all pairs of words that are semordnilaps.
    """
    if not words:
        return []
    if len(words) == 1:
        return []
    seen = set()
    result = []
    for word in words:
        if word[::-1] in seen:
            result.append((word, word[::-1]))
        seen.add(word)
    return result
