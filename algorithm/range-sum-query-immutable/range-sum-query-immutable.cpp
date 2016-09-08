class NumArray {
    vector<vector<int>> table;
public:
    NumArray(vector<int> &nums) {
        for (int i = 0; i < nums.size(); i++) {
            vector<int> tmp;
            for (int j = 0; j < i; j++) {
                tmp.push_back(0);
            }
            tmp.push_back(nums[i]);
            for (int j = i + 1; j < nums.size(); j++) {
                tmp.push_back(tmp[j - 1] + nums[j]);
            }
            table.push_back(tmp);
        }
    }

    int sumRange(int i, int j) {
        return table[i][j];
    }
};


// Your NumArray object will be instantiated and called as such:
// NumArray numArray(nums);
// numArray.sumRange(0, 1);
// numArray.sumRange(1, 2);
