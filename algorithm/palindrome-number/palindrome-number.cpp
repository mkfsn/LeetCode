#include <iostream>

using std::cout;
using std::endl;

class Solution {
public:
    bool isPalindrome(int x) {
        int y = 0, k = x;
        while (x) {
            y = (y * 10) + (x % 10);
            x /= 10;
        }
        return k >= 0 && y == k;
    }
};

int main(int argc, char* argv[]) {
    
    Solution s;
    s.isPalindrome(1001);

    return 0;
}
