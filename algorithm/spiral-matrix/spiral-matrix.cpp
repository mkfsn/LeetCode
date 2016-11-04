class Solution {
public:
    vector<int> spiralOrder(vector<vector<int>>& matrix) {
        vector<int> res;
        
        if (matrix.size() == 0 || matrix[0].size() == 0)
            return res;
        
        vector<vector<int>> direction {{0, 1}, {1, 0}, {0, -1}, {-1, 0}};
        int row = matrix.size(), col = matrix[0].size(), cursor = 0;
        vector<vector<int>> trace(row, vector<int>(col, 0));

        int x = 0, y = 0;
        for (int i = 0; i < row * col;) {
            int next_x = direction[cursor][0] + x,
                next_y = direction[cursor][1] + y;
                
            if (trace[x][y] == 0) {
                res.push_back(matrix[x][y]);
                trace[x][y] = 1;
                ++i;
            }
            
            if (next_x >= row || next_y >= col || next_x < 0 || next_y < 0 || trace[next_x][next_y] == 1) {
                cursor = (cursor + 1) % 4;
            } else {
                x = next_x;
                y = next_y;
            }

        }
        return res;
    }
};
