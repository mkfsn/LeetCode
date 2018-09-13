class Solution {
public:
    vector<int> preorder(Node* root) {
        vector<int> result;
        this->traversal(root, result);
        return result;
    }
    
    void traversal(Node* root, vector<int>& list) {
        if (root == nullptr) {
            return;
        }
        list.push_back(root->val);
        for (auto& n : root->children) {
            this->traversal(n, list);
        }
    }
};
