/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    bool isPalindrome(ListNode* head) {
        ListNode *rev, *fast, *cur;
        
        if (head == NULL || head->next == NULL)
            return true;
        
        cur = fast = head, rev = NULL;
        while (fast && fast->next) {
            cur = cur->next;
            fast = fast->next->next;
            
            head->next = rev;
            rev = head;
            head = cur;
        }
        if (fast != NULL && fast->next == NULL)
            cur = cur->next;
        
        while (rev != NULL || cur != NULL) {
            if (rev->val != cur->val)
                return false;
            rev = rev->next;
            cur = cur->next;
        }
        
        return true;
        
    }
};
