class NumArray {
    vector<int> table;
public:
    NumArray(vector<int> &nums) {
        table.push_back(0);
        for (auto n: nums) {
            table.push_back(table.back() + n);
        }
    }

    int sumRange(int i, int j) {
        return table[j + 1] - table[i];
    }
};


// Your NumArray object will be instantiated and called as such:
// NumArray numArray(nums);
// numArray.sumRange(0, 1);
// numArray.sumRange(1, 2);
