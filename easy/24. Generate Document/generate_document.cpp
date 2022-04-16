#include <string>
#include <unordered_map>

bool generateDocument(std::string characters, std::string document) {
  if (document.length() == 0) {
    return true;
  }
  std::unordered_map<char, int> charCount;
  for (char c : characters) {
    if (charCount.find(c) == charCount.end()) {
      charCount[c] = 0;
    }
    charCount[c]++;
  }
  for (int i = 0; i < document.size(); i++) {
    if (charCount.find(document[i]) == charCount.end() ||
        charCount[document[i]] == 0) {
      return false;
    }
    charCount[document[i]]--;
  }
  return true;
}
