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
    int total = 0;
public:
    int diameterOfBinaryTree(TreeNode* root) {
        height(root);
        return total;
    }
    
    int height(TreeNode* root) {
        if (root == nullptr)
            return 0;
        int left = height(root->left),
            right = height(root->right);
        total = max(total, left + right);
        return max(left, right) + 1;
    }
    
};
