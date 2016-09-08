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
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
        map<ListNode*, int> intersection;
        ListNode* tmp = headA;
        while (tmp) {
            intersection[tmp] = 0;
            tmp = tmp->next;
        }
        
        tmp = headB;
        while (tmp) {
            if (intersection.find(tmp) != intersection.end())
                return tmp;
            tmp = tmp->next;
        }
        return NULL;
    }
};
