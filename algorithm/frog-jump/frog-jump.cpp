#include <iostream>
#include <unordered_map>
#include <vector>
#include <unordered_set>

using namespace std;

class Solution {
public:
    bool canCross(vector<int>& stones) {
        unordered_map<int, unordered_set<int>> history;
        for (auto s: stones)
            history[s] = unordered_set<int>();

        history[1].insert(1);
        for (auto k = stones.begin() + 1; k != stones.end() - 1; k++) {
            for (auto h: history[*k]) {
                if (h - 1 > 0 && history.find(*k + h - 1) != history.end())
                    history[*k + h - 1].insert(h - 1);
                if (history.find(*k + h) != history.end())
                    history[*k + h].insert(h);
                if (history.find(*k + h + 1) != history.end())
                    history[*k + h + 1].insert(h + 1);
            }
        }

        return history[*(stones.end() - 1)].size() != 0;
    }
};

int main(int argc, char* argv[]) {
    vector<int> example1 = {0,1,3,5,6,8,12,17}, // true
                example2 = {0,1,2,3,4,8,9,11}; // false
    Solution s;
    cout << (s.canCross(example1) ? "True" : "False") << endl;
    cout << (s.canCross(example2) ? "True" : "False") << endl;
    return 0;
}
