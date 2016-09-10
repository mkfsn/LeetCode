class Solution {
public:
    int countPrimes(int n) {
        int end = pow(n, 0.5);
        int count = 0;
        vector<bool> prime(n, true);
        
        if (n < 2)
            return 0;
        
        prime[0] = prime[1] = false;
        for (int i = 2; i <= end; i++) {
            // If i is prime(false) or not(true)
            if (prime[i]) {
                for (int j = i + i; j < n; j += i) {
                    prime[j] = false;
                }
            }
        }
        
        for (int i = 2; i < n; i++) {
            if (prime[i])
                count++;
        }
        
        return count;
    }
};
