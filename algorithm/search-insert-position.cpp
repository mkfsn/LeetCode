class Solution {
public:
    int searchInsert(vector<int>& nums, int target) {
        for (auto it = nums.begin(); it != nums.end(); it++) {
            if (target <= *it)
                return it - nums.begin();
        }
        return nums.size();
    }
};
