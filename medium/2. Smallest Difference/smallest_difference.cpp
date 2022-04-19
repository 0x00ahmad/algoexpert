#include <algorithm>
#include <vector>

std::vector<int> smallestDifference(std::vector<int> a1, std::vector<int> a2) {
  std::vector<int> result;
  std::sort(a1.begin(), a1.end());
  std::sort(a2.begin(), a2.end());
  int i = 0;
  int j = 0;
  int min = INT_MAX;
  int min_i = 0;
  int min_j = 0;
  while (i < a1.size() && j < a2.size()) {
    int diff = a1[i] - a2[j];
    if (diff < 0) {
      diff = -diff;
    }
    if (diff < min) {
      min = diff;
      min_i = i;
      min_j = j;
    }
    if (a1[i] < a2[j]) {
      i++;
    } else {
      j++;
    }
  }
  result.push_back(a1[min_i]);
  result.push_back(a2[min_j]);
  return result;
}
