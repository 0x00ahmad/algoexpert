// This is an input class. Do not edit.
class BST {
public:
  int value;
  BST *left = nullptr;
  BST *right = nullptr;

  BST(int value) { this->value = value; }
};

struct TraversalInfo {
  int visitedNodeCount;
  int latestNodeVal;
};

void reverseInOrderFindValue(BST *tree, int k, TraversalInfo &info);

int findKthLargestValueInBst(BST *tree, int k) {
  auto info = TraversalInfo{0, -1};
  reverseInOrderFindValue(tree, k, info);
  return info.latestNodeVal;
}

void reverseInOrderFindValue(BST *tree, int k, TraversalInfo &info) {
  if (tree == nullptr || info.visitedNodeCount >= k)
    return;
  reverseInOrderFindValue(tree->right, k, info);
  if (info.visitedNodeCount < k) {
    info.visitedNodeCount++;
    info.latestNodeVal = tree->value;
    reverseInOrderFindValue(tree->left, k, info);
  }
}
