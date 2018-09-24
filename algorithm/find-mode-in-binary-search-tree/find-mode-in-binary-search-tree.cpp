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
    
    int max = 0;
    map<int, int> table;
    
    void traversal(TreeNode* root) {
        if (root == nullptr)
            return;
        
        if (table.find(root->val) == table.end())
            table[root->val] = 0;
        ++ table[root->val];
        if (table[root->val] > max)
            max = table[root->val];
        
        traversal(root->left);
        traversal(root->right);
    }
     
public:
    vector<int> findMode(TreeNode* root) {
        this->traversal(root);
        
        vector<int> result;
        for (auto const& x : this->table) {
            if (x.second == this->max) {
                result.push_back(x.first);
            }
        }
        
        return result;
    }
};
