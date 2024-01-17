#include <unordered_set>
#include <vector>
#include <math.h>

int firstDuplicateValue(std::vector<int> array) {
  for (int i = 0; i < array.size(); i++) {
    if (array[std::abs(array[i]) - 1] < 0) {
      return std::abs(array[i]);
    }
    array[std::abs(array[i]) - 1] *= -1;
  }
  return -1;
}
