from typing import List
from collections import Counter

ANAGRAMS = List[List[str]]
WORDS = List[str]

# w - words
# n - len of longest word

# O(w * n ^ 2)
def groupAnagrams(words: WORDS) -> ANAGRAMS:
    anagrams: ANAGRAMS = []
    for word in words:
        appended = False
        for idx, anagram_set in enumerate(anagrams):
            if len(anagram_set[0]) == len(word) and Counter(anagram_set[0]) == Counter(
                word
            ):
                appended = True
                anagrams[idx].append(word)
                break
        if not appended:
            anagrams.append([word])
    return anagrams


# O(w * n log(n))
def groupAnagramsDict(words: WORDS) -> ANAGRAMS:
    anagrams = {}
    for word in words:
        sorted_word = "".join(sorted(word))
        if sorted_word in anagrams:
            anagrams[sorted_word].append(word)
        else:
            anagrams[sorted_word] = [word]
    return list(anagrams.values())
