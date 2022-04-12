#include <vector>

std::vector<int> findThreeLargestNumbers(std::vector<int> array) {
  int first = INT_MIN;
  int second = INT_MIN;
  int third = INT_MIN;

  for (int i = 0; i < array.size(); i++) {
    if (array[i] > first) {
      third = second;
      second = first;
      first = array[i];
    } else if (array[i] > second) {
      third = second;
      second = array[i];
    } else if (array[i] > third) {
      third = array[i];
    }
  }

  return std::vector<int>{third, second, first};
}
