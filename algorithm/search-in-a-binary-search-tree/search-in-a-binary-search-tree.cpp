class Solution {
public:
    TreeNode* searchBST(TreeNode* root, int val) {
        queue<TreeNode*> q;
        q.push(root);
        while (!q.empty()) {
            auto n = q.front();
            q.pop();
            if (n->left != nullptr) {
                q.push(n->left);
            }
            if (n->right != nullptr) {
                q.push(n->right);    
            }
            if (n->val == val) {
                return n;
            }
        }
        return nullptr;
    }
};
