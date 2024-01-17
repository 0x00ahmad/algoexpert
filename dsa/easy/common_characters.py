from typing import List
from collections import Counter


def commonCharacters(strings: List[str]):
    counts = Counter(set(strings[0]))
    for each in strings[1:]:
        for k, v in Counter(set(each)).items():
            if k not in counts:
                counts[k] = v
                continue
            counts[k] += v
    return [k for k, v in counts.items() if v == len(strings)]


if __name__ == "__main__":
    out = commonCharacters(["hello", "ebolii", "roller"])
    print(out)
