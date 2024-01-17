# 12. Tandem Bicycle

Conditions:

- equal lengths of both arrays
- return based on the "fastest" boolean

The approach lies in how you sort the arrays, for the maxiumum you have to sort the
second array in desc to make sure you pair the highest with the lowest of the other
array. And for the miniumum you have to sort in asc to make sure you eliminate as much
of the highest values from the other array.

### Time Complexity

O(n log n)

### Space Complexity

O(1)
