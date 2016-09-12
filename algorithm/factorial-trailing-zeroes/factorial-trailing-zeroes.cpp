#include <iostream>

using namespace std;

class Solution {
public:
    int trailingZeroes(int n) {
        return n == 0 ? 0 : (n / 5) + trailingZeroes(n / 5);
    }
};

int main(int argc, char* argv[]) {
    Solution s;
    cout << s.trailingZeroes(1808548329) << endl;
    return 0;
}
