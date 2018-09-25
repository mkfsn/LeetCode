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
    int result = 0;
public:
    int longestUnivaluePath(TreeNode* root) {
        this->findLongest(root);
        return this->result;
    }
private:
    int findLongest(TreeNode* root) {
        if (root == nullptr)
            return 0;
        int l = this->findLongest(root->left),
            r = this->findLongest(root->right);
        l = (root->left && root->left->val == root->val) ? l + 1 : 0;
        r = (root->right && root->right->val == root->val) ? r + 1 : 0;
        this->result = max(this->result, l + r);
        return max(l, r);
    }
};
