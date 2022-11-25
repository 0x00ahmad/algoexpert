from collections import deque


class MinMaxStack:
    _min_stack: deque
    _max_stack: deque
    _stack: deque

    def __init__(self) -> None:
        self._stack = deque()
        self._min_stack = deque()
        self._max_stack = deque()
        self._min_stack.append(float("inf"))
        self._max_stack.append(-float("inf"))

    def peek(self) -> int:
        try:
            return self._stack[-1]
        except IndexError:
            return -1

    def pop(self):
        self._max_stack.pop()
        self._min_stack.pop()
        val = self._stack[-1]
        self._stack.pop()
        return val

    def push(self, number):
        self._stack.append(number)
        self._max_stack.append(max(self._max_stack[-1], number))
        self._min_stack.append(min(self._min_stack[-1], number))

    def getMin(self) -> int:
        return self._min_stack[-1]

    def getMax(self) -> int:
        return self._max_stack[-1]
