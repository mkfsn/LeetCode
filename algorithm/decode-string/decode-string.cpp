struct myTreeNode {
    string str;
    vector<pair<int, myTreeNode*>> children;

    myTreeNode(): str("") {}
    
    myTreeNode* insert(int n) {
        myTreeNode *tmp = new myTreeNode();
        this->children.push_back(make_pair(n, tmp));
        return tmp;
    }
};

class Solution {
    
    string decode(myTreeNode* node) {
        if (node->children.size() == 0)
            return node->str;
        
        string str("");
        for (auto it = node->children.begin(); it != node->children.end(); it++) {
            string tmp("");
            for (int i = 0; i < it->first; i++)
                tmp += decode(it->second);
            str += tmp;
        }
        
        return str + node->str;
    }
    
public:
    string decodeString(string s) {
        
        myTreeNode *root = new myTreeNode(), *tmp;
        stack<myTreeNode* > S; 
        int len(0);
        string str("");
        
        S.push(root);
        tmp = S.top();
        
        /* Build Tree */
        for (int i = 0; i < s.length(); i++) {
            
            if (s[i] >= '0' && s[i] <= '9') {
                
                if (str != "") {
                    tmp = tmp->insert(1);
                    tmp->str = str;
                    tmp = S.top();
                    str = "";
                }
                
                len = len * 10 + (s[i] - '0');
                
            } else if (s[i] == '[') {
                
                str = "";
                tmp = tmp->insert(len);
                S.push(tmp);
                len = 0;
                
            } else if (s[i] == ']') {
                
                tmp->str = str;
                str = "";
                S.pop();
                tmp = S.top();
                
            } else {
                
                str += s[i];
                
            }
        }
        
        if (!S.empty() && str != "") {
            tmp = S.top();
            tmp = tmp->insert(1);
            tmp->str = str;
        }
        
        /* Decode */
        str = decode(root);
        return str;
    }
};
