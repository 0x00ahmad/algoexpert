from typing import List


def zeroSumSubarray(nums: List[int]):
    """
    Given an array of integers, find if the array contains any subarray with 0 sum.
    """
    if not nums:
        return False
    if len(nums) == 1:
        return nums[0] == 0
    sum = 0
    seen = set()
    for num in nums:
        sum += num
        if sum == 0 or sum in seen:
            return True
        seen.add(sum)
    return False

