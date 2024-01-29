def missingNumbers(nums):
    return list(filter(lambda x: x not in nums, range(1, len(nums) + 3)))
