#include <iostream>

using std::cout;
using std::endl;
using std::string;

class Solution {
public:
    string toLowerCase(string str) {
        for (auto& c: str) {
            if (c >= 0x41 && c <= 0x5A) {
                c |= 32;
            }
        }
        return str;
    }
};

void check(bool b) {
    if (!b) {
        cout << "Wrong" << endl;
    }
}

int main() {
    Solution s;

    check("hello" == s.toLowerCase("Hello"));
    check("here" == s.toLowerCase("here"));
    check("lovely" == s.toLowerCase("LOVELY"));

    return 0;
}
