#include <iostream>
#include <vector>
#include <stack>

using std::vector;
using std::stack;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

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
public:
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int> answer;
        stack<TreeNode*> s;

        if (root == NULL)
            return answer;

        s.push(root);
        while (!s.empty()) {
            TreeNode* node = s.top();
            s.pop();
            answer.push_back(node->val);

            if (node->right != NULL)
                s.push(node->right);
            if (node->left != NULL)
                s.push(node->left);
        }
        return answer;        
    }
};

int main(int argc, char* argv[])
{
    return 0;
}
