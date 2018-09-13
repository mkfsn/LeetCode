#include <iostream>
#include <string>
#include <map>
#include <vector>

using std::string;
using std::map;
using std::vector;
using std::cout;
using std::endl;

struct Node {
    char ch;
    int freq;
    Node(char c) :ch(c), freq(0) {}

    bool operator <(Node const& rhs) {
        return this->freq < rhs.freq;
    }

    bool operator ==(Node const& rhs) {
        return this->freq == rhs.freq;
    }

    bool operator !=(Node const& rhs) {
        return this->freq != rhs.freq;
    }

    bool operator >(Node const& rhs) {
        return this->freq > rhs.freq;
    }

    bool operator >=(Node const& rhs) {
        return this->freq >= rhs.freq;
    }

    bool operator <=(Node const& rhs) {
        return this->freq <= rhs.freq;
    }
};

class Heap {
    virtual void heapify(int i) = 0;
};

class MaxHeap: public Heap {
    vector<Node> list;

    void heapify(int i) {
        int left = 2 * i + 1,
            right = 2 * i + 2;
        int largest = i;
        if (left < this->size() && this->list[left] > this->list[i]) {
            largest = left;
        }
        if (right < this->size() && this->list[right] > this->list[largest]) {
            largest = right;
        }
        if (largest != i) {
            Node tmp = this->list[i];
            this->list[i] = this->list[largest];
            this->list[largest] = tmp;
            this->heapify(largest);
        }
    }

public:
    MaxHeap(vector<Node> l): list(l) {
        for (int i = (this->size() - 1) / 2; i >= 0; --i) {
            this->heapify(i);
        }
    }

    Node pop() {
        if (this->list.size() == 0) {
            throw std::out_of_range("heap is empty");
        }
        Node n = this->list.at(0);
        this->list.erase(this->list.begin());

        for (int i = (this->size() - 1) / 2; i >= 0; --i) {
            this->heapify(i);
        }

        return n;
    }

    int size() {
        return this->list.size();
    }
};

class Solution {
public:
    string frequencySort(string s) {
        map<char, int> table;
        vector<Node> list;
        string result;

        int index = 0;

        for (auto& c : s) {
            if (table.find(c) == table.end()) {
                list.push_back(Node(c));
                table[c] = index++;
            }
            ++list[table[c]].freq;
        }

        MaxHeap heap(list);

        while (heap.size() > 0) {
            Node node = heap.pop();
            for (int i = 0; i < node.freq; i++)
                result += node.ch;
        }

        return result;
    }
};

int main() {
    Solution s;

    cout << s.frequencySort("tree") << endl;
    cout << s.frequencySort("cccaaa") << endl;
    cout << s.frequencySort("Aabb") << endl;
    cout << s.frequencySort("") << endl;

    return 0;
}
