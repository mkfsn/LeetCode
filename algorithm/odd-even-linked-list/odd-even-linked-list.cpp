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
    ListNode* oddEvenList(ListNode* head) {
        ListNode *even = new ListNode(-1),
                 *odd = new ListNode(-1),
                 *tmp_e = even,
                 *tmp_o = odd;
                 
        int even_or_odd = 0; // 0:even, 1:odd
        while (head) {
            if (even_or_odd) {
                tmp_e->next = head;
                tmp_e = head;
                head = head->next;
                tmp_e->next = NULL;
            } else {
                tmp_o->next = head;
                tmp_o = head;
                head = head->next;
                tmp_o->next = NULL;
            }
            even_or_odd ^= 1;
        }
        tmp_o->next = even->next;
        
        tmp_o = odd;
        tmp_e = even;
        odd = odd->next;
        delete tmp_e;
        delete tmp_o;
        return odd;
    }
};
