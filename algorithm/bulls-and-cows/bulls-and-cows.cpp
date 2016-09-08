class Solution {
public:
    string getHint(string secret, string guess) {
        int A = 0, B = 0;
        map<char, int> dict;
        bool passed[secret.length()] = {false};
        string s;
        stringstream ss(s);
        
        for (int i = 0; i < secret.length(); i++) {
            if (secret[i] == guess[i]) {
                A++;
                passed[i] = true;
            } else if (dict.find(secret[i]) == dict.end()) {
                dict[secret[i]] = 1;
            } else {
                dict[secret[i]]++;
            }
        }
        
        for (int i = 0; i < guess.length(); i++) {
            char s = guess[i];
            if (!passed[i] && dict.find(s) != dict.end() && dict[s] > 0) {
                B++;
                dict[s]--;
            }
        }
        
        ss << A << 'A' << B << 'B';
        return ss.str();
    }
};
