class Solution {
public:
    vector<int> intersection(vector<int>& nums1, vector<int>& nums2) {
        map<int, int> intersection;
        vector<int> result;
        
        for (auto n: nums1) {
            if (intersection.find(n) == intersection.end()) {
                intersection[n] = 1;
            }
        }
        
        for (auto n: nums2) {
            if (intersection.find(n) != intersection.end() && intersection[n] == 1) {
                intersection[n] = 2;
                result.push_back(n);
            }
        }
        
        return result;
    }
};
