class Solution {

public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        map<int, int> freq;
        vector<vector<int>> counts(nums.size() + 1);
        vector<int> res;
        
        for (auto n: nums) {
            try {
                freq[n]++;
            } catch (out_of_range) {
                freq[n] = 1;
            }
        }

        for (auto it = freq.begin(); it != freq.end(); it++)
            counts[it->second].push_back(it->first);
        
        for (auto row = counts.rbegin(); row != counts.rend() && k > 0; row++)
            for (auto col = row->begin(); col != row->end() && k > 0; col++, k--)
                res.push_back(*col);

        return res;
    }
};
