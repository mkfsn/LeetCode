#!/usr/bin/env python
# -*- coding: utf-8 -*-
__date__= ' 9 12, 2016 '
__author__= 'mkfsn'

class Solution(object):
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        
        if x < 0:
            res = int(str(abs(x))[::-1])
        else:
            res = int(str(x)[::-1])
            
        if res > (2 ** 31 - 1):
            return 0
        elif x < 0:
            return 0 - res
        else:
            return res
