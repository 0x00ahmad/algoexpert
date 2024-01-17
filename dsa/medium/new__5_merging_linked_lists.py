# This is an input class. Do not edit.
class LinkedList:
    def __init__(self, value):
        self.value = value
        self.next = None


def mergingLinkedLists(linkedListOne, linkedListTwo):
    l1 = linkedListOne
    l2 = linkedListTwo

    while l1 is not l2:
        if l1 is None:
            l1 = linkedListTwo
        else:
            l1 = l1.next
        if l2 is None:
            l2 = linkedListOne
        else:
            l2 = l2.next
