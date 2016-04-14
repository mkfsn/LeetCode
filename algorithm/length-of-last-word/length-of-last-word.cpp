#include <iostream>
#include <string>

using std::cout;
using std::endl;
using std::string;

class Solution {
public:
    int lengthOfLastWord(string s) {
        int i, start = s.length() - 1;
        while (start >= 0 && s[start] == ' ')
            start--;
        if (start < 0)
            return 0;
        for (i = start; i >= 0 && s[i] != ' '; i--)
            ;
        return start != i ? (start - i) : 0;
    }
};

int main(int argc, char* argv[]) {
    Solution s;
    cout << s.lengthOfLastWord("a abbc") << endl;
    return 0;
}
