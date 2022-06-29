# 9. Merge Overlapping Intervals

First we create a copy of the original list and sort them, then create a pointer array
for creating the new merged intervals. Then iterate over each of the sorted intervals,
and check

IF the end of the current is greater then the start of the next: then we set the current
intervals's first element to the max of the current-end with the next-end

ELSE just set the current to the next and push the new current to the new merged intervals' array

### Time Complexity

O(n log n)

### Space Complexity

O(n)
