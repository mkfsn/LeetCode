#include <stdio.h>

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

struct TreeNode* lowestCommonAncestor(struct TreeNode* root, struct TreeNode* p, struct TreeNode* q) {
    if (root == NULL)
        return NULL;
    else if (root == p || root == q)
        return root;

    struct TreeNode *left = lowestCommonAncestor(root->left, p, q),
                    *right = lowestCommonAncestor(root->right, p, q);

    if (left == NULL && right == NULL)
        return NULL;
    else if (left == NULL)
        return right;
    else if (right == NULL)
        return left;
    else
        return root;
}

int main(int argc, char* argv[])
{
    // Testing:

    return 0;
}
