#include <iostream>
#include <map>
#include <vector>

using namespace std;

class Solution {
public:
    bool canCross(vector<int>& stones) {
        
        map<int, vector<int>> history;
        for (auto s: stones) {
            vector<int> tmp;
            history[s] = tmp;
        }
        
        history[1].push_back(0);
        for (auto k = stones.begin() + 1; k != stones.end() - 1; k++) {
            for (auto h: history[*k]) {
                int last = *k - h;
                for (int i = -1; i <= 1; i++) {
                    int next = *k + last + i;
                    if (*k != next && history.find(next) != history.end() 
                        && history[next].end() == find(history[next].begin(), history[next].end(), *k)) {
                        history[next].push_back(*k);
                    }
                }
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
