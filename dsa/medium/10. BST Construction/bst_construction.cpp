#include <vector>
using namespace std;

// note that you can't remove values from a single-node tree, so calling the
// `remove` method on a single-node tree should not do anything.
class BST {
public:
  int value;
  BST *left;
  BST *right;

  BST(int val) {
    value = val;
    left = nullptr;
    right = nullptr;
  }

  BST &insert(int val) {
    if (val < value) {
      if (left == nullptr) {
        left = new BST(val);
      } else {
        left->insert(val);
      }
    } else {
      if (right == nullptr) {
        right = new BST(val);
      } else {
        right->insert(val);
      }
    }
    return *this;
  }

  bool contains(int val) {
    if (val == value) {
      return true;
    }

    if (val < value) {
      return left == nullptr ? false : left->contains(val);
    } else if (val > value) {
      return right == nullptr ? false : right->contains(val);
    }

    return false;
  }

  BST &remove(int val, BST *parent = nullptr) {
    if (val < value) {
      if (left != nullptr)
        left->remove(val, this);
    } else if (val > value) {
      if (right != nullptr)
        right->remove(val, this);
    } else {
      // handle the simple case
      if (left != nullptr && right != nullptr) {
        value = right->getSmallestFromRightSubtree();
        right->remove(value, this);
      } else if (parent == nullptr) {
        if (left != nullptr) {
          value = left->value;
          right = left->right;
          left = left->left;
        } else if (right != nullptr) {
          value = right->value;
          right = right->right;
          left = right->left;
        } else {
        }
      } else if (parent->left == this) {
        parent->left = left != nullptr ? left : right;
      } else if (parent->right == this) {
        parent->right = left != nullptr ? left : right;
      }
    }
    return *this;
  }

  int getSmallestFromRightSubtree() {
    if (left == nullptr) {
      return value;
    }

    return left->getSmallestFromRightSubtree();
  }
};
