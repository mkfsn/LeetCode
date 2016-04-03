class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        string common_prefix("");
        int maxlength = 0;
        
        for (int i = 0; i < strs.size(); i++) {
            if (strs[i].length() > maxlength)
                maxlength = strs[i].length();
        }
        
        for (int i = 0; i < maxlength; i++) {
            for (int j = 0; j < strs.size(); j++) {
                if (common_prefix.length() == i)
                    common_prefix += strs[j][i];
                else if (common_prefix[i] != strs[j][i])
                    return common_prefix.erase(common_prefix.size() - 1);
            }
        }

        return common_prefix;
        
    }
};
