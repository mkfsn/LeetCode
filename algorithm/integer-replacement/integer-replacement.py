#!/usr/bin/env python
# -*- coding: utf-8 -*-
__date__= ' 9 12, 2016 '
__author__= 'mkfsn'


class Solution(object):
    def integerReplacement(self, n):
        """
        :type n: int
        :rtype: int
        """

        if n == 2:
            return 1
        elif n == 1:
            return 0

        if (n % 2):
            plus = self.integerReplacement(n + 1)
            minus = self.integerReplacement(n - 1)
            return 1 + min([plus, minus])
        else:
            return 1 + self.integerReplacement(n / 2)

s = Solution()
print s.integerReplacement(2147483647)
