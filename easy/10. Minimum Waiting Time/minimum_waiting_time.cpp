#include <algorithm>
#include <vector>

int minimumWaitingTime(std::vector<int> queries) {
  std::sort(queries.begin(), queries.end());
  int totalWaitingTime = 0;
 int q_len = queries.size();
  for (int i = 0; i < q_len; i++) {
    totalWaitingTime += (q_len - (i + 1)) * queries[i];
  }
  return totalWaitingTime;
}
