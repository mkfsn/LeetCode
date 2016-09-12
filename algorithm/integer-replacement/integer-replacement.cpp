class Solution {
    int min(int a, int b) {
        return a <= b ? a : b;
    }
public:
    int integerReplacement(int n) {
        
        unsigned int x = n;
        
        if (n == 2)
            return 1;
        else if (n == 1)
            return 0;

        if (x % 2)
            return 1 + min(integerReplacement(x + 1), integerReplacement(x - 1));
        else
            return 1 + integerReplacement(x / 2);
    }
};
