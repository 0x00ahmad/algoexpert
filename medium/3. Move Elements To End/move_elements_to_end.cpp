#include <vector>

std::vector<int> moveElementToEnd(std::vector<int> array, int toMove) {
    int left = 0;
    int right = array.size() - 1;
    while (left < right) {
        while (array[left] != toMove && left < right) {
            left++;
        }
        while (array[right] == toMove && left < right) {
            right--;
        }
        if (left < right) {
            std::swap(array[left], array[right]);
        }
    }
    return array;
}

