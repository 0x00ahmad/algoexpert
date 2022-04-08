# 6. Find Closest Value In BST

This is a pretty simple rundown, we know that all the greater values are going to
be on the right-er side of the tree and vice versa. So we just have to calculate
the closest value to the target on each node until we exhaust the search.

For **navigating** to which side we traverse next:

if the target is less than the current value, we go left
else go right

### Time Complexity

O(log n)

### Space Complexity

O(1)
