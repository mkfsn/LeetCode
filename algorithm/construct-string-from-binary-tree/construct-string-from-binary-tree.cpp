class Solution {
public:
    string tree2str(TreeNode* t) {
        if (t == nullptr) {
            return "";
        }
        
        string result = std::to_string(t->val);
        
        if (t->left != nullptr) {
            result += "(" + this->tree2str(t->left) + ")";
        } else if (t->right != nullptr) {
            result += "()";
        }
        
        if (t->right != nullptr) {
            result += "(" + this->tree2str(t->right) + ")";
        }
        
        return result;
    }
};
