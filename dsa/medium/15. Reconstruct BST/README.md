# 15. Reconstruct BST

The approach here is a bit through but it's pretty simple, you have to recursively build
the bst and for each rercursive call you do the following:

- if the root index is equal to the passed list of values, then terminate the call
- otherwise pull the value from the valsList and check for it being out of bounds,
  - if it is then terminate
- we can create a subtree with it by recursing down and creating their subtrees by passing
  in the next set of bounds

### Time Complexity

O(n)

### Space Complexity

O(n+k)
