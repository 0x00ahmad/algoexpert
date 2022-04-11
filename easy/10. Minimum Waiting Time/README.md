# 10. Minimum Waiting Time

Now there is a bit of a trick in understanding the flow of this, but once you do it's
pretty simple; so I'm gonna try my best to explain in as less words as possible.
Here's how you approach this:

You have an array of +ve ints, you first sort the array. Now the reason you sort the
array is because this will make sure you have the shorter query times before the
longer ones.

After that you will iterate through the array and for each element at the index i,
you will multiple the current waiting time with the number of elements remaining in
the array.

You might be wondering why this works? This will work because each query preceding
the current one has to wait out the time of the queries before them in order to
execute. This is why we sort the array in before beginning, so each query has to
wait a minimum amount of time.

### Time Complexity

O(n log n)

### Space Complexity

O(1)
