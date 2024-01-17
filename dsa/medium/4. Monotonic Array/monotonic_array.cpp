#include <vector>

bool isMonotonic(std::vector<int> array) {
    int n = array.size();
    if (n <= 2) return true;
    bool is_increasing = true;
    bool is_decreasing = true;
    for (int i = 1; i < n; i++) {
        if (array[i] > array[i - 1]) {
            is_decreasing = false;
        } else if (array[i] < array[i - 1]) {
            is_increasing = false;
        }
    }
    return is_increasing || is_decreasing;
}
