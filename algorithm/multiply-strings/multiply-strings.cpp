class Solution {
public:
    string multiply(string num1, string num2) {
        vector<int> multiply(num1.size() + num2.size() + 1, 0);
        string res;
        int i, j;
        
        reverse(num1.begin(), num1.end());
        reverse(num2.begin(), num2.end());
        
        for (i = 0; i < num1.size(); i++) {
            for (j = 0; j < num2.size(); j++) {
                multiply[i + j] += ((int)(num1[i] - '0') * (int)(num2[j] - '0'));
            }
        }
        
        for (i = 0; i < multiply.size(); i++) {
            if (multiply[i] >= 10) {
                multiply[i + 1] += (multiply[i] / 10);
                multiply[i] %= 10;
            }
            res += (char)(multiply[i] + '0');
        }
        
        for (i = multiply.size() - 1; i >= 0 && multiply[i] == 0; i--) {
            res[i] = '\0';
        }
        res.resize(i + 1);
        reverse(res.begin(), res.end());
        
        if (res.size() == 0)
            return string("0");
        return res;
    }
};
