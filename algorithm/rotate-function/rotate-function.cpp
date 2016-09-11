class Solution {
public:
    int maxRotateFunction(vector<int>& A) {
        
        int fn = 0, max = 0, len = A.size();
        
        int sum = 0;
        for (int i = 0; i < len; i++) {
            sum += A[i];
            fn += i * A[i];
        }
        max = fn;
        
        for (int i = 1; i < len; i++) {
            fn = fn + sum - len * A[len - i];
            if (fn > max)
                max = fn;
        }
        
        return max;
    }
};
