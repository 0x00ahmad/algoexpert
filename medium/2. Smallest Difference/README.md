# 2. Smallest Difference

The approach here is to utilize sort and 2 pointers to traverse the lists side by side, and
compare out the difference. If it's zero then we found it, else increment the shorter of the
two numbers. Repeat until either of the lists is exhausted.

### Time Complexity

O(n log n)

### Space Complexity

O(1)
