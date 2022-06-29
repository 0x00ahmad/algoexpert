#include <algorithm>
#include <vector>

std::vector<std::vector<int>>
mergeOverlappingIntervals(std::vector<std::vector<int>> intervals) {
  std::vector<std::vector<int>> sortedIntervals = intervals;
  std::sort(sortedIntervals.begin(), sortedIntervals.end(),
            [](std::vector<int> a, std::vector<int> b) { return a[0] < b[0]; });

  std::vector<std::vector<int> *> mergedIntervals;
  std::vector<int> *currentInterval = &sortedIntervals[0];
  mergedIntervals.push_back(currentInterval);

  for (auto &nextInterval : sortedIntervals) {
    int currentIntervalEnd = currentInterval->at(1);
    int nextIntervalStart = nextInterval.at(0);
    int nextIntervalEnd = nextInterval.at(1);

    if (currentIntervalEnd >= nextIntervalStart) {
      currentInterval->at(1) = std::max(currentIntervalEnd, nextIntervalEnd);
    } else {
      currentInterval = &nextInterval;
      mergedIntervals.push_back(currentInterval);
    }
  }

  std::vector<std::vector<int>> mergedIntervalsVector;
  for (auto &interval : mergedIntervals) {
    mergedIntervalsVector.push_back(*interval);
  }
  return mergedIntervalsVector;
}
