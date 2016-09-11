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
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        map<int, ListNode*> path;
        ListNode *tmp = head;
        int idx = 0;
        
        while (tmp) {
            path[idx++] = tmp;
            tmp = tmp->next;
        }
        
        tmp = path[idx - n];
        if (tmp == head) {
            head = head->next;
        } else {
            path[idx - n - 1]->next = tmp->next;
        }
        
        return head;
    }
};
