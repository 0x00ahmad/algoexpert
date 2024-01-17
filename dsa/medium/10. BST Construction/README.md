# 10. BST Construction

#### Insert

This is a straightforward one, you just have to

#### Contains

A binary search down to the target node

#### Remove

This is a tad bit more complicated, first we find the target value. If the target is present,
then we find the next best value for the node we are removing to re-balance the tree

To do that, we can use the left most leaf node of the right sub-tree of the target node.

### Time Complexity

Insert: O(log n)
Contains: O(log n)
Remove: O(log n)

### Space Complexity

O(log n) for all methods

**Note:** Worst case scenario for all the complexities above is O(n)
