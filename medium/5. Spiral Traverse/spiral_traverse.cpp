#include <vector>

std::vector<int> spiralTraverse(std::vector<std::vector<int>> array) {
    std::vector<int> result;
    int row = 0;
    int col = 0;
    int row_end = array.size() - 1;
    int col_end = array[0].size() - 1;
    int direction = 0;
    while (row <= row_end && col <= col_end) {
        if (direction == 0) {
            for (int i = col; i <= col_end; i++) {
                result.push_back(array[row][i]);
            }
            row++;
            direction = 1;
        } else if (direction == 1) {
            for (int i = row; i <= row_end; i++) {
                result.push_back(array[i][col_end]);
            }
            col_end--;
            direction = 2;
        } else if (direction == 2) {
            for (int i = col_end; i >= col; i--) {
                result.push_back(array[row_end][i]);
            }
            row_end--;
            direction = 3;
        } else if (direction == 3) {
            for (int i = row_end; i >= row; i--) {
                result.push_back(array[i][col]);
            }
            col++;
            direction = 0;
        }
    }
    return result;
}

