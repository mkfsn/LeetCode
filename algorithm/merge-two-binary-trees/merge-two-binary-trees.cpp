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
    TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
        if (t1 == nullptr && t2 == nullptr) {
            return nullptr;
        } else if (t1 == nullptr) {
            TreeNode* t = new TreeNode(t2->val);
            t->left = t2->left;
            t->right = t2->right;
            return t;
        } else if (t2 == nullptr) {
            TreeNode* t = new TreeNode(t1->val);
            t->left = t1->left;
            t->right = t1->right;
            return t;
        }
        TreeNode *t = new TreeNode(t1->val + t2->val);
        t->left = this->mergeTrees(t1->left, t2->left);
        t->right = this->mergeTrees(t1->right, t2->right);
        return t;
    }
};
