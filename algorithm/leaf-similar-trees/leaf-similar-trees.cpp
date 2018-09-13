class Solution {
public:
    bool leafSimilar(TreeNode* root1, TreeNode* root2) {
        vector<int> left, right;
        
        this->traversal(root1, left);
        this->traversal(root2, right);
        
        return left == right;
    }
    
    void traversal(TreeNode* node, vector<int>& leaves) {
        if (node == nullptr) {
            return; 
        } else if (node->left == nullptr && node->right == nullptr) {
            leaves.push_back(node->val);
        }
        this->traversal(node->left, leaves);
        this->traversal(node->right, leaves);
    }
};
