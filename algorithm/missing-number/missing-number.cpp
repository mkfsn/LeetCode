class Solution {
public:
    int missingNumber(vector<int>& nums) {
        int sum = 0, N = nums.size();
        
        for (auto n: nums)
            sum += n;
        
        return ((N * (N + 1)) / 2) - sum;
    }
};
