# 24. Generate Document

Approach: have a hashmap with counts of the characters and iterate over the document,
and if any count drops below 0, it's not a valid document.

### Time Complexity

O(n+m)

### Space Complexity

O(c)
