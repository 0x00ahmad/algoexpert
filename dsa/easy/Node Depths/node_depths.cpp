#include <iostream>

class BinaryTree {
public:
  int value;
  BinaryTree *left;
  BinaryTree *right;

  BinaryTree(int value) {
    this->value = value;
    left = nullptr;
    right = nullptr;
  }
};

int recurse(BinaryTree *node, int &level) {
  if (node == nullptr) {
    return 0;
  }
  int nl = level + 1;
  int l = recurse(node->left, nl);
  int r = recurse(node->right, nl);

  return l + r + level;
}

int nodeDepths(BinaryTree *root) {
  if (root == nullptr) {
    return 0;
  }
  int level= 0;
  int out = recurse(root, level);
  return out;
}

int main() { return 0; }
