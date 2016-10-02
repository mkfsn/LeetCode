#include<iostream>
#include<vector>

using namespace std;

class Solution {
    
public:
    int coinChange(vector<int>& coins, int amount) {
        vector<int> dp(amount + 1, amount + 1);

        dp[0] = 0;
        for (int i = 1; i <= amount; ++i) {
            for (int j = 0; j < coins.size(); ++j) {
                if (coins[j] <= i) {
                    dp[i] = min(dp[i], dp[i - coins[j]] + 1);
                }
            }
        }

        return dp[amount] > amount ? -1 : dp[amount];
    }
};

int main(int argc, char* argv[]) {

    vector<int> coins = {10, 2};
    int amount = 3;

    Solution s;

    cout << s.coinChange(coins, amount) << endl;

    return 0;
}
