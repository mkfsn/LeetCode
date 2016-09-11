class Solution {
public:
    bool containsNearbyDuplicate(vector<int>& nums, int k) {
        map<int, int> indexes;
        for (auto it = nums.begin(); it != nums.end(); it++) {
            if (indexes.find(*it) != indexes.end()) {
                if (it - nums.begin() - indexes[*it] <= k)
                    return true;
            }
            indexes[*it] = (it - nums.begin());
        }
        return false;
    }
};
