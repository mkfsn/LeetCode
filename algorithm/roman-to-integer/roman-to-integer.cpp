#include <iostream>
#include <map>
#include <string>

using std::map;
using std::string;

class Solution {
public:
    int romanToInt(string s) {
        int sum = 0, prev = 0;

        map<char, int> hash;
        hash['I'] = 1;
        hash['V'] = 5;
        hash['X'] = 10;
        hash['L'] = 50;
        hash['C'] = 100;
        hash['D'] = 500;
        hash['M'] = 1000;

        for (int i = 0; i < s.length(); i++) {
            if (prev < hash[s[i]])
                sum -= prev * 2;
            sum += int(hash[s[i]]);
            prev = hash[s[i]];
        }

        return sum;
    }
};

int main(int argc, char* argv[])
{
    Solution s;

    std::cout << s.romanToInt("XXIX");

    return 0;
}
