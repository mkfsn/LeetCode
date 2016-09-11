class Solution {
public:
    vector<int> plusOne(vector<int>& digits) {
        int carry = 0;
        digits.back() += 1;
        for (auto it = digits.rbegin(); it != digits.rend(); it++) {
            *it += carry;
            if (*it >= 10) {
                *it -= 10;
                carry = 1;
            } else {
                carry = 0;
            }
        }
        if (carry != 0)
            digits.insert(digits.begin(), carry);
        return digits;
    }
};
