#include <vector>

std::vector<int> insertionSort(std::vector<int> array) {
  for (int i = 1; i < array.size(); i++) {
    int j = i;
    while (j > 0 && array[j] < array[j - 1]) {
      int temp = array[j];
      array[j] = array[j - 1];
      array[j - 1] = temp;
      j--;
    }
  }
  return array;
}
