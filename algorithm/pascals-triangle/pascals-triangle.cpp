class Solution {
public:
    vector<vector<int>> generate(int numRows) {
        vector<vector<int>> result;
        
        if (numRows == 0)
            return result;
        
        vector<int> v1 = {1};
        result.push_back(v1);
        
        if (numRows == 1)
            return result;
        
        vector<int> prev = {1, 1};
        result.push_back(prev);
        
        if (numRows == 2)
            return result;
        
        vector<int> tmp;
        for (int i = 2; i < numRows; i++) {
            tmp.clear();
            tmp.push_back(1);
            for (int j = 1; j < prev.size(); j++) {
                tmp.push_back(prev[j] + prev[j - 1]);
            }
            tmp.push_back(1);
            result.push_back(tmp);
            prev = tmp;
        }
        return result;
    }
};
