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
    
    vector<TreeNode*> DFS(TreeNode* root, TreeNode* node) {
        stack<TreeNode*> stk;
        vector<TreeNode*> path;
        TreeNode* tmp;
        
        stk.push(root);
        
        while (!stk.empty() && stk.top()) {
            tmp = stk.top();
            stk.pop();
            
            
            while (path.size() > 0) {
                TreeNode* last = path.back();
                if (last->left == tmp || last->right == tmp)
                    break;
                path.pop_back();
            }
            path.push_back(tmp);
            
            if (tmp == node)
                break;
            
            if (tmp->left != NULL)
                stk.push(tmp->left);
            if (tmp->right != NULL)
                stk.push(tmp->right);
        }
        
        return path;
    }
    
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        
        vector<TreeNode*> path_p = DFS(root, p),
                        path_q = DFS(root, q);
        map<TreeNode*, bool> path;
        
        for (auto n: path_p) {
            path[n] = true;
        }
        
        for (auto it = path_q.rbegin(); it != path_q.rend(); it++) {
            if (path.find(*it) != path.end())
                return *it;
        }
        
        return NULL;
    }
};
