#!/usr/bin/env python
# -*- coding: utf-8 -*-
__date__ = ' 3 30, 2016 '
__author__ = 'mkfsn'

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    # @param {ListNode} l1
    # @param {ListNode} l2
    # @return {ListNode}
    def mergeTwoLists(self, l1, l2):
        head = tail = ListNode(0)
        
        while l1 and l2:
            if l1.val < l2.val:
                node = ListNode(l1.val)
                head.next = node
                head = node
                l1 = l1.next
            else:
                node = ListNode(l2.val)
                head.next = node
                head = node
                l2 = l2.next
        
        if l1 is None:
            while l2:
                node = ListNode(l2.val)
                head.next = node
                head = node
                l2 = l2.next
        else:
            while l1:
                node = ListNode(l1.val)
                head.next = node
                head = node
                l1 = l1.next
        return tail.next
