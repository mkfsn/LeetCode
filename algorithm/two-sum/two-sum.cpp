class Solution {
    
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<int> ans;
        for (vector<int>::iterator it = nums.begin(); it !=  nums.end(); it++) {
            vector<int>::iterator jt = find(it + 1, nums.end(), target - *it);
            if (jt != nums.end()) {
                ans.push_back(it - nums.begin());
                ans.push_back(jt - nums.begin());
                return ans;
            }
        }
        return ans;
    }
};
