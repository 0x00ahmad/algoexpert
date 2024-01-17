#include <algorithm>
#include <iostream>
#include <vector>

std::vector<int> sortedSquaredArray(std::vector<int> &A) {
  std::vector<int> result;
  for (int i = 0; i < A.size(); i++) {
    result.push_back(A[i] * A[i]);
  }
  std::sort(result.begin(), result.end());
  return result;
}

int main() { return 0; }
