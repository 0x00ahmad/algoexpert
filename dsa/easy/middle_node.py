# This is an input class. Do not edit.
class LinkedList:
    def __init__(self, value):
        self.value = value
        self.next = None


def middleNode(linkedList: LinkedList):
    if linkedList is None:
        return None
    count = 0
    head = linkedList
    while head is not None:
        head = head.next
        count += 1
    mid = count // 2
    head = linkedList
    while mid > 0:
        head = head.next
        mid -= 1
    return head
