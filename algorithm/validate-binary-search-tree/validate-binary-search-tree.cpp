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
    
    TreeNode* findNode(TreeNode* root, TreeNode* node) {
        TreeNode* tmp = root;
        while (tmp) {
            if (tmp->val == node->val)
                return tmp;
            else if (tmp->val < node->val)
                tmp = tmp->right;
            else if (tmp->val > node->val)
                tmp = tmp->left;
        }
        return tmp;
    }
public:
    bool isValidBST(TreeNode* root) {

        vector<TreeNode*> nodes;
        queue<TreeNode*> BFS;
        TreeNode *tmp;
        
        BFS.push(root);
        while(!BFS.empty() && BFS.front()) {
            tmp = BFS.front();
            BFS.pop();
            if (tmp->left) {
                BFS.push(tmp->left);
                nodes.push_back(tmp->left);
            }
            if (tmp->right) {
                BFS.push(tmp->right);
                nodes.push_back(tmp->right);
            }
        }
        
        for (auto n: nodes) {
            if (n != findNode(root, n))
                return false;
        }
        return true;
    }
};
