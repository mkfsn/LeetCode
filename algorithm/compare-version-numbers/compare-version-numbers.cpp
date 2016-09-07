class Solution {

    vector<string> split(const string &source, const char &delim) {
        stringstream ss(source);
        string token;
        vector<string> res;
        while (getline(ss, token, delim))
            res.push_back(token);
        return res;
    }

public:


    int compareVersion(string version1, string version2) {

        if (version1 == version2)
            return 0;

        vector<string> ver1 = split(version1, '.');
        vector<string> ver2 = split(version2, '.');
        
        int len = ver1.size();
        if (ver2.size() > len)
            len = ver2.size();
        
        for (int i = 0; i < len; i++) {
            int v1, v2;
            if (i >= ver1.size())
                v1 = 0;
            else
                v1 = atoi(ver1[i].c_str());

            if (i >= ver2.size())
                v2 = 0;
            else
                v2 = atoi(ver2[i].c_str());
            
            if (v1 < v2)
                return -1;
            else if (v1 > v2)
                return 1;
        }

        return 0;
    }
};
