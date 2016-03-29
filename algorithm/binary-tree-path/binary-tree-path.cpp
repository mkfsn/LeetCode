#include <iostream>
#include <string>
#include <vector>
#include <stack>
#include <sstream>

using std::stack;
using std::vector;
using std::string;
using std::pair;
using std::make_pair;
using std::stringstream;
using std::cout;
using std::endl;

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
    vector<string> binaryTreePaths(TreeNode* root) {

        vector<string> answer;
        vector<int> path;
        stringstream ss;
        stack<pair<TreeNode*, int>> s;
        pair<TreeNode*, int> t;
        TreeNode *node = NULL;
        int level = 1;

        if (root == NULL)
            return answer;

        s.push(make_pair(root, level));
        while (!s.empty()) {
            t = s.top();
            s.pop();

            node = t.first;
            level = t.second;
            if (path.size() >= level)
                path[level - 1] = node->val;
            else
                path.push_back(node->val);

            if (node->left == NULL && node->right == NULL) {
                string str("");
                for (int i = 0; i < level - 1; i++) {
                    ss << path[i];
                    str += (ss.str() + "->");
                    ss.str("");
                }
                ss << path[level - 1];
                str += ss.str();
                ss.str("");
                answer.push_back(str);
            }
            
            if (node->right != NULL)
                s.push(make_pair(node->right, level + 1));
            if (node->left != NULL)
                s.push(make_pair(node->left, level + 1));
        }
        return answer;
    }
};

TreeNode* make_node(int val)
{
    TreeNode *node = new TreeNode(val);
    return node;
}

int main(int argc, char* argv[])
{

    TreeNode *root = make_node(1);
    root->left = make_node(2);
    root->right = make_node(3);
    root->left->left = make_node(4);

    Solution s;
    vector<string> v;

    v = s.binaryTreePaths(root);
    for (int i = 0; i < v.size(); i++)
        cout << v[i] << endl;

    return 0;
}
