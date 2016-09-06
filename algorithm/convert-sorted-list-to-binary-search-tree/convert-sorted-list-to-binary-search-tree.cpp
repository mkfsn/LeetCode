/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
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

    ListNode* buildTree(TreeNode* root, ListNode* head) {
        if (root->left == NULL && root->right == NULL) {
            root->val = head->val;
            return head->next;
        }
        
        if (root->left != NULL) {
            head = buildTree(root->left, head);
        }
        
        root->val = head->val;
        head = head->next;
        
        if (root->right != NULL) {
            head = buildTree(root->right, head);
        }
        
        return head;
    }

    TreeNode* sortedListToBST(ListNode* head) {
        if (head == NULL)
            return NULL;
        
        TreeNode *root, *cur;
        ListNode *L;
        root = new TreeNode(0);
        
        /* Use BFS to build empty tree */
        queue<TreeNode*> Q;
        Q.push(root);
        L = head->next;
        
        while (L != NULL) {
            cur = Q.front();
            Q.pop();
            
            if (L != NULL) {
                cur->left = new TreeNode(0);
                Q.push(cur->left);
                L = L->next;
            }
            
            if (L != NULL) {
                cur->right = new TreeNode(0);
                Q.push(cur->right);
                L = L->next;
            }
        }
        buildTree(root, head);
        return root;
    }
};
