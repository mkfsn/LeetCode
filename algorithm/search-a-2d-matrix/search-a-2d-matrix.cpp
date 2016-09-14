class Solution {
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        if (matrix.size() == 0)
            return false;
        
        int size = matrix.size() * matrix[0].size();
        int left = 0, right = size - 1, middle, i, j, x;
        
        while (left <= right) {
            middle = (left + right) / 2;
            i = middle / matrix[0].size();
            j = middle % matrix[0].size();
            x = matrix[i][j];
            
            if (x < target) {
                left = middle + 1;
            } else if (x > target) {
                right = middle - 1;
            } else {
                return true;
            }
        }
        return false;
    }
};
