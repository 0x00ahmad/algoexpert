#include <string>

std::string caesarCypherEncryptor(std::string str, int key) {
  std::string result = "";
  for (int i = 0; i < str.length(); i++) {
    if (str[i] >= 'a' && str[i] <= 'z') {
      result += (str[i] + key - 'a') % 26 + 'a';
    } else if (str[i] >= 'A' && str[i] <= 'Z') {
      result += (str[i] + key - 'A') % 26 + 'A';
    } else {
      result += str[i];
    }
  }
  return result;
}

