class Solution {
public:
    vector<vector<int>> generateMatrix(int n) {
        vector<vector<int>> direction {{0, 1}, {1, 0}, {0, -1}, {-1, 0}};
        vector<vector<int>> res(n, vector<int>(n, 0));
        int cursor = 0;
        int x = 0, y = 0;
        
        for (int i = 1; i <= n * n; ) {
            int next_x = direction[cursor][0] + x,
                next_y = direction[cursor][1] + y;
            
            if (res[x][y] == 0)
                res[x][y] = i++;
            
            if (next_x >= n || next_y >= n || next_x < 0 || next_y < 0 || res[next_x][next_y] != 0)
                cursor = (cursor + 1) % 4;
            else
                x = next_x, y = next_y;
        }
        
        return res;
    }
};
