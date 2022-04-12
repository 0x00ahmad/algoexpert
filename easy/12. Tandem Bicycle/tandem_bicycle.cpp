#include <algorithm>
#include <bits/stdc++.h>
#include <vector>

int tandemBicycle(std::vector<int> redShirtSpeeds,
                  std::vector<int> blueShirtSpeeds, bool fastest) {
  int result = 0;
  std::sort(redShirtSpeeds.begin(), redShirtSpeeds.end());
  if (fastest) {
    std::sort(blueShirtSpeeds.begin(), blueShirtSpeeds.end(),
              std::greater<int>());
  } else {
std::sort(blueShirtSpeeds.begin(), blueShirtSpeeds.end(),
              std::greater<int>());

  }

  for (int i = 0; i < redShirtSpeeds.size(); i++) {
    result += std::max(redShirtSpeeds[i], blueShirtSpeeds[i]);
  }

  return result;
}
