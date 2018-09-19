class Solution {
public:
    vector<vector<int>> flipAndInvertImage(vector<vector<int>>& A) {
        for (auto& row: A) {
            int l = 0, r = row.size() - 1;
            while (l <= r) {
                int tmp = row[l];
                row[l] = row[r] ^ 1;
                row[r] = tmp ^ 1;
                ++l, --r;
            }
        }
        return A;
    }
};

