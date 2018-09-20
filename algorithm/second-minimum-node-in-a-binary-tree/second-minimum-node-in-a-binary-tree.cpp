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
    vector<int> values;
    
    void dfs(TreeNode* root) {
        if (root == nullptr) return;
        this->dfs(root->left);
        this->dfs(root->right);
        this->update(root->val);
    }
    
    void update(int v) {
        if (this->values.size() == 0) {
            this->values.push_back(v);
            return;
        } else if (this->values.size() == 1) {
            if (v == this->values[0])
                return;
            if (v > this->values[0]) {
                this->values.push_back(v);
                return;                
            } else {
                this->values.insert(this->values.begin(), v);
                return;
            }
        }
        if (v == this->values[0] || v == this->values[1]) {
            return;            
        }
        
        if (v > this->values[0] && v < this->values[1]) {
            this->values.pop_back();
            this->values.push_back(v);
        } else if (v < this->values[0]) {
            this->values.pop_back();
            this->values.insert(this->values.begin(), v);
        }
    }
public:
    int findSecondMinimumValue(TreeNode* root) {
        this->dfs(root);
        if (this->values.size() < 2)
            return -1;
        return this->values[1];
    }
    
};
