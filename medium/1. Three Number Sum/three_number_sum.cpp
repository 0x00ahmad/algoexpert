#include <algorithm>
#include <vector>

// The first typical three sum optimal solution
std::vector<std::vector<int>> threeNumberSum(std::vector<int> array,
                                             int targetSum) {
  std::vector<std::vector<int>> result;
  std::sort(array.begin(), array.end());

  for (int i = 0; i < array.size(); i++) {
    int left = i + 1;
    int right = array.size() - 1;

    while (left < right) {
      int sum = array[i] + array[left] + array[right];

      if (sum == targetSum) {
        result.push_back({array[i], array[left], array[right]});
        left++;
        right--;
      } else if (sum < targetSum) {
        left++;
      } else {
        right--;
      }
    }
  }
  return result;
}
