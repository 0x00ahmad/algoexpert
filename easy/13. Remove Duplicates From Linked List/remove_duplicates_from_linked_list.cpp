// This is an input struct. Do not edit.
class LinkedList {
public:
  int value;
  LinkedList *next = nullptr;

  LinkedList(int value) { this->value = value; }
};

LinkedList *removeDuplicatesFromLinkedList(LinkedList *linkedList) {
  LinkedList *head = linkedList;
  LinkedList *current = linkedList;
  while (current != nullptr) {
    LinkedList *next = current->next;
    while (next != nullptr && next->value == current->value) {
      next = next->next;
    }
    current->next = next;
    current = next;
  }
  return head;
}
