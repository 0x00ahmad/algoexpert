#include <vector>

class BST {
public:
  int value;
  BST *left;
  BST *right;

  BST(int value) {
    this->value = value;
    left = nullptr;
    right = nullptr;
  }

  void insert(int value) {
    if (value < this->value) {
      if (left == nullptr) {
        left = new BST(value);
      } else {
        left->insert(value);
      }
    } else {
      if (right == nullptr) {
        right = new BST(value);
      } else {
        right->insert(value);
      }
    }
  }
};

// start the root for the tree from the middle of the vector
BST *minHeightBst(std::vector<int> array) {
  if (array.size() == 0) {
    return nullptr;
  }

  int mid = array.size() / 2;
  BST *root = new BST(array[mid]);

  std::vector<int> left(array.begin(), array.begin() + mid);
  std::vector<int> right(array.begin() + mid + 1, array.end());

  root->left = minHeightBst(left);
  root->right = minHeightBst(right);

  return root;
}
