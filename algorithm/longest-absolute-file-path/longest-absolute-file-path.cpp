#include <iostream>
#include <vector>
#include <stack>

using namespace std;

class Solution {
public:
    int lengthLongestPath(string input) {
        stack<int> path;
        int length = 0, max_length = 0;
        
        for (int i = 0; i < input.length();) {
            int level = 1;
            bool hasdot = false;
            string tmp;
            
            while (input[i] == '\t') {
                ++level;
                ++i;
            }

            while (i < input.length() && input[i] != '\n') {
                tmp += input[i];
                if (input[i] == '.')
                    hasdot = true;
                ++i;
            }

#if DEBUG
            for (int i = 1; i < level; i++)
                cout << '\t';
            cout << tmp;
            cout << '(' << tmp.length() << ')' << endl;
#endif

            while (i < input.length() && input[i] == '\n')
                ++i;

            if (level > path.size()) {
                // Sub dir
                path.push(tmp.length());
                length += tmp.length();
            } else if (level < path.size()) {
                // Parent dir
                for (int i = path.size(); i >= level; --i) {
                    length -= path.top();
                    path.pop();
                }
                path.push(tmp.length());
                length += tmp.length();
            } else {
                // Next dir
                length -= path.top();
                path.pop();
                path.push(tmp.length());
                length += tmp.length();
            }

            if (hasdot) {
                // Is file
                if (length > max_length)
                    max_length = length;
            } else {
                ++path.top();
                ++length;
            }
        }
        
        return max_length;
    }
};

int main(int argc, char* argv[]) {
    Solution s;
    cout << s.lengthLongestPath("dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext") << endl;
    cout << s.lengthLongestPath("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext") << endl;
    cout << s.lengthLongestPath("dir\n    file.txt") << endl;
    return 0;
}
