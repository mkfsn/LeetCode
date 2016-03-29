#include <iostream>
#include <vector>
#include <set>

using std::set;
using std::vector;
using std::pair;
using std::make_pair;

class Solution {
public:
    vector<int> findMinHeightTrees(int n, vector<pair<int, int>>& edges) {

        vector<int> leaves;
        if (n == 1) {
            leaves.push_back(0);
            return leaves;
        }

        vector<set<int>> adj(n);
        for (int i = 0; i < edges.size(); i++) {
            int x = edges[i].first,
                y = edges[i].second;
            adj[x].insert(y);
            adj[y].insert(x);
        }

        for (int i = 0; i < n; i++)
            if (adj[i].size() == 1)
                leaves.push_back(i);

        while (n > 2) {
            n -= leaves.size();
            vector<int> newleaves;
            for (auto i: leaves) {
                set<int>::iterator it = adj[i].begin();
                int j = *it;
                adj[i].erase(j);
                adj[j].erase(i);
                if (adj[j].size() == 1)
                    newleaves.push_back(j);
            }
            leaves = newleaves;
        }

        return leaves;
    }
};

int main(int argc, char* argv[])
{
    vector<pair<int, int>> edges;
    edges.push_back(make_pair(1, 0));
    edges.push_back(make_pair(1, 2));
    edges.push_back(make_pair(1, 3));
    Solution s;
    s.findMinHeightTrees(4, edges);
    return 0;
}
