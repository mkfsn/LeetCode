class TrieNode {
    map<char, TrieNode*> children;
public:
    // Initialize your data structure here.
    TrieNode() {
        
    }
    
    TrieNode* find(char c) {
        if (this->children.find(c) == this->children.end())
            return NULL;
        return this->children[c];
    }
    
    TrieNode* insert(char c) {
        return (this->children[c] = new TrieNode());
    }
};

class Trie {
public:
    Trie() {
        root = new TrieNode();
    }

    // Inserts a word into the trie.
    void insert(string word) {
        TrieNode *cur = root, *tmp;
        for (auto c: word) {
            if ((tmp = cur->find(c)) == NULL) {
                tmp = cur->insert(c);
            }
            cur = tmp;
        }
        cur->insert('\0');
    }

    // Returns if the word is in the trie.
    bool search(string word) {
        TrieNode *cur = root, *tmp;
        for (auto c: word) {
            if ((tmp = cur->find(c)) == NULL)
                return false;
            cur = tmp;
        }
        return cur->find('\0') != NULL;
    }

    // Returns if there is any word in the trie
    // that starts with the given prefix.
    bool startsWith(string prefix) {
        TrieNode *cur = root, *tmp;
        for (auto c: prefix) {
            if ((tmp = cur->find(c)) == NULL)
                return false;
            cur = tmp;
        }
        return true;
    }

private:
    TrieNode* root;
};

// Your Trie object will be instantiated and called as such:
// Trie trie;
// trie.insert("somestring");
// trie.search("key");
