# 6. Longest Peak

With each iteration, first we find the target number which is a peak (with numbers on it's
either sides being smaller) and then we expand both of them upto the first false condition.
Then check and set the max of the current peak with the longest ongoing peak.

### Time Complexity

O(n)

### Space Complexity

O(1)
