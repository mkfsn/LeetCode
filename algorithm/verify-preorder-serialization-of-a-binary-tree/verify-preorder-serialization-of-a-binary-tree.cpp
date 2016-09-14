class Solution {
    
    vector<string> split(string str, const char delim) {
        stringstream ss(str);
        string target;
        vector<string> res;
        
        while (getline(ss, target, delim)) {
            res.push_back(target);
        }
        
        return res;
    }
    
public:
    bool isValidSerialization(string preorder) {
        vector<string> serialized = split(preorder, ',');
        stack<string> Q;
        int i = 0;
        
        if (serialized.size() == 0)
            return true;
        
        if (serialized[i] != "#")
            Q.push(serialized[i]);
        i++;
        while (!Q.empty() && i < serialized.size()) {
            Q.pop();
            
            if (i >= serialized.size()) {
                return false;
            } else if (serialized[i] != "#") {
                Q.push(serialized[i]);
            }
            i++;
            
            if (i >= serialized.size()) {
                return false;
            } else if (serialized[i] != "#") {
                Q.push(serialized[i]);
            }
            i++;
        }
        
        return Q.empty() && i == serialized.size();
    }
};
