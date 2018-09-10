#include <iostream>
#include <vector>
#include <set>

using std::cout;
using std::endl;
using std::string;
using std::vector;
using std::set;

class Solution {
    vector<string> morsecode;
public:
    Solution() :morsecode{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."} {}

    int uniqueMorseRepresentations(vector<string>& words) {
        set<string> table;    

        for (auto& word: words) {
            string code;

            for (auto& c: word) {
                code += this->morsecode[c - 'a'];
            }

            table.insert(code);
        }


        return table.size();
    }
};

int main() {
    Solution s;
    vector<string> input {"gin", "zen", "gig", "msg"};

    cout << (2 == s.uniqueMorseRepresentations(input) ? "Correct" : "Incorrect") << endl;

    return 0;
}
