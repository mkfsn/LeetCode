class Solution {
public:
    int maxDepth(Node* root) {
        if (root == nullptr) {
            return 0;
        }
        
        int max = 0;
        for (auto& node : root->children) {
            auto n = this->maxDepth(node);
            if (n > max)
                max = n;
        }
        return 1 + max;
    }
};
