class Solution {
    int biSearch(vector<int>& nums, int n, int target) {
        int middle, left = 0, right = n - 1;
        while (left <= right) {
            middle = (left + right) / 2;
            if (target > nums[middle])
                left = middle + 1;
            else if (target < nums[middle])
                right = middle - 1;
            else
                return middle;
        }
        return -1;
    }
public:
    vector<int> twoSum(vector<int>& numbers, int target) {
        int idx;
        vector<int> res;

        for (int i = numbers.size() - 1; i >= 0; i--) {
            if ((idx = biSearch(numbers, i, target - numbers[i])) != -1) {
                res.push_back(idx + 1);
                res.push_back(i + 1);
            }
        }
        return res;
    }
};
