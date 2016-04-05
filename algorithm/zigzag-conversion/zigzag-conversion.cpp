#include <iostream>
#include <vector>
#include <string>

using std::cout;
using std::endl;
using std::string;
using std::vector;

class Solution {
public:
    string convert(string s, int numRows) {
        vector<string> rows(numRows);
        string ans;
        
        if (numRows == 1)
            return s;
        
        for (int i = 0; i < s.length();) {
            for (int j = 0; j < numRows && i < s.length(); i++, j++) {
                rows[j] += s[i];
            }
            for (int j = numRows - 2; j > 0 && i < s.length(); i++, j--) {
                rows[j] += s[i];
            }
        }
        
        for (int i = 0; i < numRows; i++) {
            for (int j = 0; j < rows[i].size(); j++) {
                ans += rows[i][j];
            }
        }
        
        return ans;
    }
};

int main(int argc, char* argv[]) {
    Solution s;
    cout << s.convert("olsydisqprnyxwmokuxrndpdvmncntlbeebhotyfmnfjjsytox", 3);
    return 0;
}
