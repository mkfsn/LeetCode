class Solution {
public:
    vector<double> averageOfLevels(TreeNode* root) {
        vector<double> result;
        
        if (root == nullptr) {
            return result;
        }
        
        queue<TreeNode*> q;
        q.push(root);
        
        while (!q.empty()) {
            
            vector<TreeNode*> nodes;
            double sum = 0;
            int total = 0;
            while (!q.empty()) {
                auto n = q.front();
                nodes.push_back(n);
                sum += n->val;
                ++total;    
                q.pop();
            }
            result.push_back(sum / total);
            
            for (auto& node: nodes) {
                if (node->left)
                    q.push(node->left);
                if (node->right)
                    q.push(node->right);
            }
        }
        return result;
        
    }
};
