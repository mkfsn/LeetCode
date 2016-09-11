#include <iostream>

using std::cout;
using std::endl;

class Solution {
public:
    int reverse(int x) {

        bool negative = x < 0;
        if (negative)
            x = -x;
            
        int y = 0;
        int tmp;
        
        while (x) {
            tmp = y;
            y = (y * 10) + (x % 10);
            if ((y - (x % 10)) / 10 != tmp)
                return 0;
            x /= 10;
        }
        
        if (y > 0x7fffffff || y < 0)
            return 0;
        
        if (negative)
            return -y;
        else
            return y;
    }
};

int main(int argc, char* argv[]) {
    Solution s;

    cout << s.reverse(1534236469) << endl;

    return 0;
}
