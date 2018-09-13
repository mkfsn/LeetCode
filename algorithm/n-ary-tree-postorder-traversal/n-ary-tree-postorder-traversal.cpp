class Solution {
public:
    vector<int> postorder(Node* root) {
        vector<int> result;
        this->traversal(root, result);
        return result;
    }
    
    void traversal(Node* root, vector<int>& list) {
        if (root == nullptr) {
            return;
        }
        for (auto& n : root->children) {
            this->traversal(n, list);
        }
        list.push_back(root->val);
    }
};
