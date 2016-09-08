class Solution {
    char isAlphaAndConvert(char c) {
        if (c >= 'a' && c <= 'z')
            return c;
        else if (c >= 'A' && c <= 'Z')
            return c - 'A' + 'a';
        else if (c >= '0' && c <= '9')
            return c;
        else
            return '\0';
    }
public:
    bool isPalindrome(string s) {
        char left, right;
        for (int i = 0, j = s.length() - 1; i < j; i++, j--) {
            while (i < s.length() && '\0' == (left = isAlphaAndConvert(s[i])))
                i++;
            while (j >= 0 && '\0' == (right = isAlphaAndConvert(s[j])))
                j--;
            if (left != right)
                return false;
        }
        return true;
    }
};
