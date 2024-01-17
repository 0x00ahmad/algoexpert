#include <iostream>

class BST {
public:
  int value;
  BST *left;
  BST *right;

  BST(int val);
  BST &insert(int val);
};

int findClosestValueInBst(BST *tree, int target) {
  int current_closest = tree->value;
  BST *current = tree;
  while (current != NULL) {
    // this means that the current value is closer to the target
    // than our current closest value
    if (std::abs(current->value - target) <
        std::abs(current_closest - target)) {
      current_closest = current->value;
    }
    // if the target is less than the current value, we go left
    if (current->value > target) {
      current = current->left;
    } else {
      current = current->right;
    }
  }
  return current_closest;
}

int main() { return 0; }
