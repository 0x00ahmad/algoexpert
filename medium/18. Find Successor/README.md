# 18. Find Successor

The slower one is to run once and record the in order traversal in an array, and run
through that node again and return it's successor

The faster one involves only searching for a valid candidate that is:
- the parent of the left most leaf of the right most subtree (from the left child)

The key takeaway is to go upto the parent of the node(which is the left child of that
parent we visited)

### Time Complexity

O(n)

### Space Complexity

O(n)
