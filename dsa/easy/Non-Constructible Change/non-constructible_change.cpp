#include <algorithm>
#include <vector>

int nonConstructibleChange(std::vector<int> coins) {
  int c_size = coins.size();
  if (c_size == 0)
    return 1;
  std::sort(coins.begin(), coins.end());
  if (coins[0] > 1)
    return 2;
  int curr_change = 0;
  for (int i = 0; i < c_size; i++) {
    if (coins[i] > curr_change + 1) {
      return curr_change + 1;
    }
    curr_change += coins[i];
  }
  return curr_change + 1;
}
