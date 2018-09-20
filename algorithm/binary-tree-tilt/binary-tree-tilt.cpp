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
    int sum = 0;
public:
    int findTilt(TreeNode* root) {
        this->dfs(root);
        return sum;
    }
    
    int dfs(TreeNode* root) {
        if (root == nullptr)
            return 0;
        int left = this->dfs(root->left),
            right = this->dfs(root->right);
        sum += std::abs(left - right);
        return left + right + root->val;
    }
};
