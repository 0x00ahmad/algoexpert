# 5. Non-Constructible Change

The approach is a bit mathematical but I'll try my best to explain.

So the way you can think of this is that there are n coins in the array (ascending order),
and to find the minimum non-constructable change is to find the first coin that is
greater than the prev_sum + 1 (where prev sum is the sum of all previous coins up until the current)
OR the whole array summed + 1.

Why this works is that we can take the following example:

Array = [1 1 2 7]
sum = 0

iteration#1 : sum = 1, elem = 1 ~ sum(0) + 1 = 1> elem(1) ? yes
iteration#2 : sum = 2, elem = 1 ~ sum(1) + 1 = 2 > elem(1) ? yes
iteration#3 : sum = 4, elem = 2 ~ sum(4) + 1 = 5 > elem(2) ? yes
Break here as the current sum 4 + 1 is less than 7

If you look closely, and think of this, we can construct all changes 1,2,3,4 with the given coins,
but not 5 which hence; is the minimum non-constructable change

You can take more examples and draw them out to check this but hope this explaination is clear enough :]

### Time Complexity

O(n log(n)) ~ the sort time

### Space Complexity

O(1)
