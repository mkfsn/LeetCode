class Solution {
public:
    vector<int> getRow(int rowIndex) {
        vector<int> result;
        
        result.push_back(1);
        if (rowIndex == 0)
            return result;
        
        result.push_back(1);
        if (rowIndex == 1)
            return result;
        
        vector<int> tmp;
        for (int i = 2; i <= rowIndex; i++) {
            tmp.clear();
            tmp.push_back(1);
            for (int j = 1; j < result.size(); j++) {
                tmp.push_back(result[j] + result[j - 1]);
            }
            tmp.push_back(1);
            result = tmp;
        }
        return tmp;        
    }
};
