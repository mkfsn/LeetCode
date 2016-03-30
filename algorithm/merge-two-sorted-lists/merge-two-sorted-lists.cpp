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
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
        ListNode *rp = new ListNode(0),
                 *hp = rp;
        
        while (l1 && l2) {
            if (l1->val < l2->val) {
                hp->next = l1;
                l1 = l1->next;
            } else {
                hp->next = l2;
                l2 = l2->next;
            }
            hp = hp->next;
        }
        
        if (l1)
            hp->next = l1;
        if (l2)
            hp->next = l2;
        
        hp = rp;
        rp = rp->next;
        delete hp;
        return rp;
    }
};

