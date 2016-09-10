class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        map<int, int> counts;
        for (auto it = nums.begin(); it != nums.end();) {
            if (counts.find(*it) == counts.end())
                counts[*it] = 0;
            
            counts[*it]++;
            
            if (counts[*it] > 2)
                it = nums.erase(it);
            else
                it++;
        }
        return nums.size();
    }
};
