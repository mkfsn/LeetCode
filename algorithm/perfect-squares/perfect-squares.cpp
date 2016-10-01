#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    int numSquares(int n) {
        static vector<int> dp{0};
        for (int i = dp.size(); i <= n; ++i) {
            int squares = INT_MAX;
            for (int j = 1; j * j <= i; ++j)
                squares = min(squares, dp[i - j * j] + 1);                            
            dp.push_back(squares);
        }
        return dp[n];
    }
};

int main(int argc, char* argv[]) {
    Solution s;
    cout << s.numSquares(8829) << endl;
    return 0;
}
