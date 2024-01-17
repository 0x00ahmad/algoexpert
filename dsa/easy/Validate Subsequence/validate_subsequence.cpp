#include <iostream>
#include <vector>

bool isValidSubsequence(std::vector<int> array, std::vector<int> sequence) {
  int i = 0;
  for (int j = 0; j < array.size(); j++) {
    if (i == sequence.size()) break;
    if (array[j] == sequence[i]) {
      i++;
    }
  }
  return i == sequence.size();
}

int main() { return 0; }
