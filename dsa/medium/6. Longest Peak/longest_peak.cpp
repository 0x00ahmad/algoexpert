#include <vector>

int longestPeak(std::vector<int> array) {
  int longestPeakLength = 0;
  int i = 1;
  while (i < int(array.size() - 1)) {
    if (!(array[i - 1] < array[i] && array[i + 1] < array[i])) {
      i++;
      continue;
    }

    int leftIdx = i - 2;
    while (leftIdx >= 0 && array[leftIdx] < array[leftIdx + 1]) {
      leftIdx--;
    }

    int rightIdx = i + 2;
    while (rightIdx < int(array.size()) &&
           array[rightIdx] < array[rightIdx - 1]) {
      rightIdx++;
    }

    int currentPeakLength = rightIdx - leftIdx - 1;
    longestPeakLength = std::max(longestPeakLength, currentPeakLength);
    i = rightIdx;
  }
  return longestPeakLength;
}
