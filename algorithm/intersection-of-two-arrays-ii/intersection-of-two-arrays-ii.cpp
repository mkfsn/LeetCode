class Solution {
public:
    vector<int> intersect(vector<int>& nums1, vector<int>& nums2) {
        map<int, int> intersection;
        vector<int> result;
        for (auto n: nums1) {
            if (intersection.find(n) == intersection.end())
                intersection[n] = 0;
            intersection[n]++;
        }
        
        for (auto n: nums2) {
            if (intersection.find(n) != intersection.end() && intersection[n] > 0) {
                result.push_back(n);
                intersection[n]--;
            }
        }
        
        return result;
    }
};
