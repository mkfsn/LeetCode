#!/usr/bin/env python
# -*- coding: utf-8 -*-
__date__= ' 9 02, 2016 '
__author__= 'mkfsn'


class Solution(object):
    def moveZeroes(self, nums):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        zeros = nums.count(0)
        for i in range(zeros):
            nums.remove(0)
            nums.append(0)
