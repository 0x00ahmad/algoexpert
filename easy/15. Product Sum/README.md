# 15. Product Sum

You have an array which can have potentially nested arrays. The solution is simply:

For array [x, y] = x + y
For array [x, [y,z]] = x + 2 \* (y + z)
For array [x, [y, [z]]] = x + 2 \*(y+ 3z)
For array [x, ...] = x + 1 \* (next + n \* ...)

You have to sum the arrays up but essentially recursively call each child array to
sum it's children up.

### Time Complexity

O(n)

### Space Complexity

O(d)
