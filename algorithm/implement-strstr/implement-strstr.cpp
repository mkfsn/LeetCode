class Solution {
public:
    int strStr(string haystack, string needle) {
        int i = 0;
        
        while (i + needle.length() <= haystack.length()) {
            
            if (0 == haystack.compare(i, needle.length(), needle)) {
                return i;
            }
            
            i++;
        }
        
        return -1;
    }
};
