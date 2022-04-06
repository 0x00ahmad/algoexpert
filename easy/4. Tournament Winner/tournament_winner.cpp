#include <iostream>
#include <string>
#include <unordered_map>
#include <vector>

std::string tournamentWinner(std::vector<std::vector<std::string>> competitions,
                             std::vector<int> results) {
  std::unordered_map<std::string, int> scores;
  for (int i = 0; i < competitions.size(); i++) {
    std::string curr_winner = competitions[i][results[i] == 1 ? 0 : 1];
    if (scores.find(curr_winner) == scores.end()) {
      scores[curr_winner] = 0;
    }
    scores[curr_winner] += 3;
  }
  std::string winner = "";
  int max = 0;
  for (auto it = scores.begin(); it != scores.end(); it++) {
    // only if it's greater than the current max
    if (it->second > max) {
      max = it->second;
      winner = it->first;
    }
  }
  return winner;
}
