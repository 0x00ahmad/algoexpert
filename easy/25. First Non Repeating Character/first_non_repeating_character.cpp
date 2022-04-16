#include <string>
#include <unordered_map>

int firstNonRepeatingCharacter(std::string s) {
  std::unordered_map<char, int> charCount;
  for (int i = 0; i < s.length(); i++) {
    charCount[s[i]]++;
  }
  for (int i = 0; i < s.length(); i++) {
    if (charCount[s[i]] == 1) {
      return i;
    }
  }
  return -1;
}
