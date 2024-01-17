#include <iostream>
#include <map>
#include <vector>

std::vector<int> twoNumberSum(std::vector<int> array, int targetSum) {
  std::map<int, int> dict;
  for (int i = 0; i < array.size(); i++) {
    int num = array[i];
    int diff = targetSum - num;
    if (dict.find(diff) != dict.end()) {
      return std::vector<int> {diff,num};
    }
    dict[num] = i;
  }
  return {};
}

int main() { return 0; }
