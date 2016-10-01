#include <iostream>
#include <vector>

using namespace std;

class Solution {
    int findMinPair(vector<int> counts, int start, int end) {
        int i = 1, j, min = counts[i] + counts[target - i];
        for (i = 2, j = i * i; j < target; i++, j = i * i) {
            int tmp = counts[j] + counts[target - j];
            if (tmp < min)
                min = tmp;
        }
        return min;
    }
public:
    int numSquares(int n) {

        vector<int> counts(n + 1, 0);

        // Init
        for (int i = 1, j = i * i; j <= n; i++, j = i * i)
            counts[j] = 1;

        counts[2] = 2;
        for (int i = 3; i <= n; i++) {
            if (counts[i] == 1)
                continue;
            counts[i] = findMinPair(counts, 1, i - 1);
        }
        
        return counts[n];
    }
};

int main(int argc, char* argv[]) {
    Solution s;
    cout << s.numSquares(8829) << endl;
    return 0;
}
