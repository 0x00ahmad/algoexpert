#include <algorithm>
#include <bits/stdc++.h>
#include <vector>

bool classPhotos(std::vector<int> redShirtHeights,
                 std::vector<int> blueShirtHeights) {
  std::sort(redShirtHeights.begin(), redShirtHeights.end(),
            std::greater<int>());
  std::sort(blueShirtHeights.begin(), blueShirtHeights.end(),
            std::greater<int>());

  bool colorsInFirstRow = redShirtHeights[0] < blueShirtHeights[0];
  for (int i = 0; i < redShirtHeights.size(); i++) {
    int red_height = redShirtHeights[i];
    int blue_height = blueShirtHeights[i];
    if (colorsInFirstRow) {
      if (red_height >= blue_height) {
        return false;
      }
    } else {
      if (blue_height >= red_height) {
        return false;
      }
    }
  }
  return true;
}
