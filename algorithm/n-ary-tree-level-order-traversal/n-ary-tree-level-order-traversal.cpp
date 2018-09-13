class Solution {
public:
    vector<vector<int>> levelOrder(Node* root) {
        
        vector<vector<int>> result;
        
        if (root == nullptr) {
            return result;
        }
        
        queue<Node*> q;
        q.push(root);
        
        while (!q.empty()) {
            vector<Node*> nodes;
            vector<int> level;
            
            while (!q.empty()) {
                auto n = q.front();
                nodes.push_back(n);
                level.push_back(n->val);
                q.pop();
            }
            result.push_back(level);
            
            for (auto& node: nodes) {
                for (auto& n: node->children) {
                    q.push(n);
                }
            }
        }
        
        return result;
    }
};
