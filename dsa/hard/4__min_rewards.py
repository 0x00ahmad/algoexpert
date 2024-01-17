from typing import List


def minRewards(scores: List[int]) -> int:
    rewards = [1 for _ in scores]
    # first from left to right
    for i in range(1, len(scores)):
        if scores[i] > scores[i - 1]:
            rewards[i] = rewards[i - 1] + 1
    # then from right to left
    for i in reversed(range(len(scores) - 1)):
        if scores[i] > scores[i + 1]:
            rewards[i] = max(rewards[i], rewards[i + 1] + 1)
    return sum(rewards)
