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
    vector<TreeNode*> nodes;
    
    void traversal(TreeNode *root) {
        if (root == nullptr)
            return;
        this->traversal(root->left);
        this->nodes.push_back(root);
        this->traversal(root->right);
    }
    
    void assemble() {
        TreeNode* prev = nullptr;
        for (auto& n: this->nodes) {
            if (prev != nullptr) {
                prev->right = n;
            }
            prev = n;
            n->left = n->right = nullptr;
        }
    }
    
public:
    TreeNode* increasingBST(TreeNode* root) {
        this->traversal(root);
        this->assemble();
        if (this->nodes.size() == 0)
            return nullptr;
        return this->nodes[0];
    }
};
