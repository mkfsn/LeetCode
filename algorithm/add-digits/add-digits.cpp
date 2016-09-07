class Solution {
public:
    int addDigits(int num) {
        if (num < 10)
            return num;
        int m = num % 9;
        return m == 0 ? 9 : m;
    }
};
