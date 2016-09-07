class Solution {
    vector<string> split(const string& source, const char& delim) {
        stringstream ss(source);
        string token;
        vector<string> result;
        while (getline(ss, token, delim)) {
            result.push_back(token);
        }
        return result;
    }
public:
    bool wordPattern(string pattern, string str) {
        vector<string> strs = split(str, ' ');
        map<string, char> dict1;
        map<char, string> dict2;
        
        for (int i = 0; i < pattern.length() || i < strs.size(); i++) {
            
            if (i >= pattern.length() || i >= strs.size())
                return false;

            char c = pattern[i];
            string s = strs[i];
            
            if (dict1.find(s) == dict1.end()) {
                dict1[s] = c;
            }
            if (dict2.find(c) == dict2.end()) {
                dict2[c] = s;
            }
            
            if (dict1[s] != c || dict2[c] != s) {
                return false;
            }
        }
        return true;
    }
};
