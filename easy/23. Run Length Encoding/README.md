# 23. Run Length Encoding

This is one of those count up till the first invalid match found problems.

Iterate over the string and keep a count of the running character till you see the
next character.

As for appending to the result, if the count > 9, then append upto 9 and reduce the
count and append again.

#### Example

running character = A
count = 17

result = "3C"

After 1st append:

count = 8
result = "3C9A"

After 2nd append:

count = 1
result = "3C9A8A"

### Time Complexity

O(n)

### Space Complexity

O(1)
