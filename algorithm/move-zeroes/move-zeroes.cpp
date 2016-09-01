class Solution {
public:
    void moveZeroes(vector<int>& nums) {
        int count = 0;
        for (auto it = nums.begin(); it != nums.end(); ) {
            if (*it == 0) {
                it = nums.erase(it);
                count++;
            } else {
                ++it;
            }
        }
        while (count--) {
            nums.push_back(0);
        }
    }
};
