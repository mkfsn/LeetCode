#include <iostream>
#include <vector>

using std::cout;
using std::endl;
using std::vector;

class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        if (nums.size() < 2)
            return nums.size();
        
        int n = nums[0];
        for (auto it = nums.begin() + 1; it != nums.end();) {
            if (*it == n) {
                it = nums.erase(it);
            } else {
                n = *it;
                it++;
            }
        }

        return nums.size();
    }
};

int main(int argc, char* argv[]) {
    Solution s;
    vector<int> nums = {1, 1};
    s.removeDuplicates(nums);
    return 0;
}
