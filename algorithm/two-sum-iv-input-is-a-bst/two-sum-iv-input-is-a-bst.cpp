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
    set<int> s;
public:
    bool findTarget(TreeNode* root, int k) {
        if (root == nullptr) {
            return false;
        }
        if (s.count(k - root->val) != 0) {
            return true;
        }
        s.insert(root->val);
        return this->findTarget(root->left, k) || this->findTarget(root->right, k);
    }
};
