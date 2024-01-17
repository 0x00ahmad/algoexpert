#include <iostream>
#include <vector>

// This is the class of the input root. Do not edit it.
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

// don't forget to pass sums by reference :^]
void recurse(BinaryTree *node, int runningSum, std::vector<int> &sums) {
  if (node == nullptr) {
    return;
  }

  runningSum += node->value;

  if (node->left == nullptr && node->right == nullptr) {
    sums.push_back(runningSum);
    return;
  }

  recurse(node->left, runningSum, sums);
  recurse(node->right, runningSum, sums);
}

std::vector<int> branchSums(BinaryTree *root) {
  std::vector<int> result = {};
  if (root == nullptr) {
    return result;
  }
  recurse(root, 0, result);
  return result;
}
