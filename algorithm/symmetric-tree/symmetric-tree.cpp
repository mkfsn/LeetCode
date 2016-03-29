#include <iostream>

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
    bool isMirror(TreeNode *left, TreeNode *right) {
        if (left == NULL && right == NULL)
            return true;
        else if (left == NULL || right == NULL)
            return false;
        
        bool l = isMirror(left->left, right->right),
             r = isMirror(left->right, right->left);
        return l && r && left->val == right->val;
    }
public:
    bool isSymmetric(TreeNode* root) {
        if (root == NULL)
            return true;
        if (root->left == NULL && root->right == NULL)
            return true;
        else if (root->left == NULL || root->right == NULL)
            return false;
        return isMirror(root->left, root->right);
    }
};

int main(int argc, char* argv[])
{
    return 0;
}
