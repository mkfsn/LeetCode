class Solution {
public:
    bool isPerfectSquare(int num) {
        for (int i = 3; num > 1 ; i+=2) {
            num -= i;
        }
        return num == 1;
    }
};
