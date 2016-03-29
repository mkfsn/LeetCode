#include <iostream>
#include <vector>
#include <queue>

using std::cout;
using std::endl;
using std::vector;
using std::queue;
using std::make_pair;
using std::pair;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;

    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    vector<vector<int>> levelOrder(TreeNode* root) {

        vector<vector<int>> answer;
        queue<pair<TreeNode*, int>> q;
        pair<TreeNode*, int> t;
        int level = 1;

        if (root == NULL)
            return answer;

        q.push(make_pair(root, level));
        answer.resize(level);

        while (!q.empty()) {

            t = q.front();
            q.pop();
            if (level != t.second) {
                answer.resize(t.second);
            }

            level = t.second;
            answer[level - 1].push_back(t.first->val);

            if (t.first->left != NULL)
                q.push(make_pair(t.first->left, level + 1));
            if (t.first->right != NULL)
                q.push(make_pair(t.first->right, level + 1));
        }
        return answer;
    }
};

int main(int argc, char* argv[])
{
    return 0;
}
