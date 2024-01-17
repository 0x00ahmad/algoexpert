# 17. Binary Tree Diameter

If you think about it in an abstract way, the diameter is just the depth of the left and right
sub-trees, so to find it we can just find the depth of each subtree and return the max(
  each subtree's left&right children's depths,
  the diameter of the left&right children
)

### Time Complexity

O(n)

### Space Complexity

O(1)
