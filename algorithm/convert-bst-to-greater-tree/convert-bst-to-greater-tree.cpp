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
    int sum;
public:
    Solution() :sum(0) {}
    TreeNode* convertBST(TreeNode* root) {
        if (root == nullptr)
            return nullptr;
        convertBST(root->right);
        root->val += this->sum;
        this->sum = root->val;
        convertBST(root->left);
        return root;
    }
};
