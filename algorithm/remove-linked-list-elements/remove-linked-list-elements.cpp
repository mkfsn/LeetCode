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
    ListNode* removeElements(ListNode* head, int val) {
        ListNode *prev, *cur;
        while (head && head->val == val) {
            prev = head;
            head = head->next;
            delete prev;
        }
        prev = head;
        if (prev)
            cur = prev->next;
        else
            cur = NULL;
        
        while (cur) {
            if (cur->val == val) {
                prev->next = cur->next;
                delete cur;
                cur = prev->next;
            } else {
                prev = cur;
                cur = cur->next;
            }
        }
        return head;
    }
};
