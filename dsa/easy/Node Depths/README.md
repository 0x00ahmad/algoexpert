# 8. Node Depths

Pretty simple, perform a recursive traversal and on each child node call, pass the
current depth of the parent + 1 to for the next child node's depth. And return the
sum of = left sub-tree + right sub-tree + current depth level.

### Time Complexity

O(n)

### Space Complexity

O(1) ~ excluding the call stack space
