# 8. First Duplicate Value

The 1st solution I came up with used a set to keep track of the found elements.

The solution I settled with was the one with constant space, which is basically using the
fact that the numbers don't cross 'n' which is the length of the array so if there are
duplicate values in there we can find it as we will revisit the same index twice with those
duplicate values. We can keep track by negating the values at that index.

### Time Complexity

O(n)

### Space Complexity

O(n)
