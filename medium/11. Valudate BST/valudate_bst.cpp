#include <iostream>

class BST {
public:
  int value;
  BST *left;
  BST *right;

  BST(int val);
  BST &insert(int val);
};

// minimum signed int value
int MIN = -2147483648;
int MAX = 2147483647;

bool _validateBst(BST *tree, int min, int max);

bool validateBst(BST *tree) { return _validateBst(tree, MIN, MAX); }

bool _validateBst(BST *tree, int min, int max) {
  if (tree->value < min || tree->value >= max) {
    return false;
  }
  if (tree->left != NULL && !_validateBst(tree->left, min, tree->value)) {
    return false;
  }
  if (tree->right != NULL && !_validateBst(tree->right, tree->value, max)) {
    return false;
  }
  return true;
}
