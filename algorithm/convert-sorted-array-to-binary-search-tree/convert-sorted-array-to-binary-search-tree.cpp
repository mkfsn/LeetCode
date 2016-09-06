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

    TreeNode* buildTree(vector<int>& nums, int l, int r) {
        if (l > r)
            return NULL;
        
        int m = (l + r) / 2;
        TreeNode *root = new TreeNode(nums[m]);
        
        if (l == r)
            return root;
        
        root->left = buildTree(nums, l, m - 1);
        root->right = buildTree(nums, m + 1, r);
        return root;
    }

    TreeNode* sortedArrayToBST(vector<int>& nums) {
        return buildTree(nums, 0, nums.size() - 1);
    }
};
