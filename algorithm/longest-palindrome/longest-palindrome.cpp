class Solution {
public:
    int longestPalindrome(string s) {
        map<char, int> M;
        int sum = 0, other = 0;
        
        for (auto c: s) {
            try {
                ++M[c];
            } catch(...) {
                M[c] = 1;
            }
        }
        
        for (auto it = M.begin(); it != M.end(); it++) {
            if (it->second >= 2) {
                other += it->second & 1;
                sum += ((it->second >> 1) << 1);
            } else {
                other += it->second;
            }
        }
        
        return other > 0 ? sum + 1 : sum;
    }
};
