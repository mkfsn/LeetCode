#include <iostream>
#include <vector>

using std::cout;
using std::endl;
using std::vector;

class Solution {
public:
    vector<int> countBits(int num) {
        vector<int> bits(num + 1, 0);
        for (int i = 1, j = i << 1, k = i + 1; i <= num; j <<= 1) {
            bits[i] = 1;
            for (; k <= num && k < j; k++)
                bits[k] = bits[i] + bits[k - i];
            i = k;
        }
        return bits;
    }
};

int main(int argc, char* argv[]) {
    Solution s;
    vector<int> result = s.countBits(16);

    for (auto n: result)
        cout << n << ", ";
    cout << endl;

    return 0;
}
