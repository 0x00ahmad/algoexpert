# 5. Spiral Traverse

Since there are 4 directions in a matrix, we can use a "switch" to keep track of what
direction we are currently travelling in, and with each iteration go from point "a" to "b"
in the direction, in order of right, down, left, up. And at the end of each, nudge the
row/col count by one to keep spiraling down to the center.

### Time Complexity

O(n)

### Space Complexity

O(m)
