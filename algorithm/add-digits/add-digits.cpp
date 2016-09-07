class Solution {
public:
    int addDigits(int num) {
        if (num < 10)
            return num;
        int m = 9 + num % 9;
        return m % 10 + m / 10;
    }
};
