#include <string>

void numerificate(std::string &result, char c, int count) {
  if (count > 9) {
    while (count > 9) {
      result += std::to_string(9) + c;
      count -= 9;
    }
  }
  result += std::to_string(count) + c;
}

std::string runLengthEncoding(std::string str) {
  std::string result = "";
  int count = 1;
  for (int i = 1; i < str.length(); i++) {
    if (str[i] == str[i - 1]) {
      count++;
    } else {
      numerificate(result, str[i - 1], count);
  count = 1;
    }
  }
  numerificate(result, str[str.length() - 1], count);
  return result;
}
