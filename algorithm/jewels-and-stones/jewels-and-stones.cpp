#include <iostream>
#include <map>

using std::cout;
using std::endl;
using std::map;
using std::string;

class Solution {
public:
    int numJewelsInStones(string J, string S) {
        map<char, int> table;
        for (auto c: S) {
            if (table.find(c) == table.end()) {
                table[c] = 0;
            }
            table[c]++;
        }

        int result = 0;
        for (auto c: J) {
            if (table.find(c) != table.end()) {
                result += table[c];
            }
        }
        return result;
    }
};

int main() {

    Solution s;

    if (3 != s.numJewelsInStones("aA", "aAAbbbb") ) {
        cout << "Wrong" << endl;
    }

    if (0 != s.numJewelsInStones("z", "ZZ") ) {
        cout << "Wrong" << endl;
    }
    
    return 0;
}
