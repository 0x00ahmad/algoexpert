# 7. Branch Sums

The approach is fairly simple, recursively call the child nodes until a leaf is
reached, at which point you have to sum up the leaf node's value with the branche's
values. We do this by keeping track of a 'running sum' where you will pass down the
sum of the previous + current node value.

## Time Complexity

O(n)

### Space Complexity

O(n) ~ excluding the space accounted for the recursive call stack.
