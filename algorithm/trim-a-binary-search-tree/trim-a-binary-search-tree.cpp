class Solution {
public:
    TreeNode* trimBST(TreeNode* root, int L, int R) {
        
        if (root == nullptr) {
            return nullptr;
        }
        
        if (root->val < L) {
            return this->trimBST(root->right, L, R);
        }
        
        if (root->val > R) {
            return this->trimBST(root->left, L, R);
        }
        
        root->left = this->trimBST(root->left, L, R);
        root->right = this->trimBST(root->right, L, R);
        return root;
    }
};
