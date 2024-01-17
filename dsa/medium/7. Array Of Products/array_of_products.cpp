#include <vector>

std::vector<int> arrayOfProducts(std::vector<int> array) {
  std::vector<int> result(array.size());
  for (int i = 0; i < array.size(); i++) {
    int product = 1;
    for (int j = 0; j < array.size(); j++) {
      if (i != j) {
        product *= array[j];
      }
    }
    result[i] = product;
  }
  return result;
}
