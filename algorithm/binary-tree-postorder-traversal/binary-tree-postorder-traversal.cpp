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
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int> answer;
        stack<TreeNode*> s, output;
        TreeNode *node;

        if (root == NULL)
            return answer;

        s.push(root);
        while (!s.empty()) {
            node = s.top();
            output.push(node);
            s.pop();
            if (node->left != NULL)
                s.push(node->left);
            if (node->right != NULL)
                s.push(node->right);
        }

        while (!output.empty()) {
            node = output.top();
            output.pop();
            answer.push_back(node->val);
        }
        return answer;        
    }
};

int main(int argc, char* argv[])
{
    return 0;
}
