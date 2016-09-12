#include <iostream>

class Solution {
public:
    vector<vector<int>> levelOrderBottom(TreeNode* root) {
        vector<vector<int>> result;
        if (root == NULL)
            return result;

        queue<TreeNode*> Q;
        queue<int> D;

        Q.push(root);
        D.push(0);

        while (!Q.empty()) {
            TreeNode* node = Q.front();
            int depth = D.front();

            Q.pop();
            D.pop();

            if (node->left) {
                Q.push(node->left);
                D.push(depth + 1);
            }
            if (node->right) {
                Q.push(node->right);
                D.push(depth + 1);
            }

            if (result.size() < depth + 1) {
                vector<int> tmp;
                result.push_back(tmp);
            }

            result[depth].push_back(node->val);
        }

        reverse(result.begin(), result.end());
        return result;
    }
};
