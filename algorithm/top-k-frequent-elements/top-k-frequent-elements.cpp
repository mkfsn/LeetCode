class Solution {
    
    static bool compare(pair<int, int> const& i, pair<int, int> const& j) {
        return i.second > j.second;
    }

public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        map<int, int> freq;
        vector<int> res;
        
        for (int i = 0; i < nums.size(); i++) {
            if (freq.find(nums[i]) == freq.end())
                freq[nums[i]] = 0;
            freq[nums[i]]++;
        }
        
        vector<pair<int,int>> vp(freq.begin(), freq.end());
        sort(vp.begin(), vp.end(), compare);
        
        for (auto it = vp.begin(); it != vp.end() && it - vp.begin() < k; it++)
            res.push_back(it->first);
            
        return res;
    }
};
