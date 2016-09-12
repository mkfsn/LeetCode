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
    int countNodes(TreeNode* root) {
        if (root == NULL)
            return 0;
        
        TreeNode *left, *right;
        int max_height = 0;
        int count;
        
        left = right = root->left;
        while (right) {
            left = left->left;
            right = right->right;
            ++max_height;
        }
        count = (1 << max_height) - 1;
        
        if (left == NULL)
            return 1 + count + countNodes(root->right);
            
        return 1 + count + countNodes(root->left);
    }
};
