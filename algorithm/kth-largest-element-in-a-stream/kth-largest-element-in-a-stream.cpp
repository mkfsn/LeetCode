#include <iostream>
#include <vector>

using std::vector;
using std::string;
using std::cout;
using std::endl;

class KthLargest {

    int capacity;
    vector<int> numbers;

    void insert(int val) {
        if (this->numbers.size() < this->capacity) {
            this->numbers.push_back(val);
        } else if (this->numbers.at(0) < val) {
            this->numbers.erase(this->numbers.begin());
            this->numbers.push_back(val);
        }

        for (int i = this->numbers.size() / 2; i >= 0; --i) {
            this->min_heapify(i);
        }
    }

    void min_heapify(int index) {
        int left = index * 2 + 1,
            right = index * 2 + 2;
        int smallest_node = index;

        if (left < this->numbers.size() && this->numbers[left] < this->numbers[index] ) {
            smallest_node = left;
        }

        if (right < this->numbers.size() && this->numbers[right] < this->numbers[smallest_node]) {
            smallest_node = right;
        }

        if (smallest_node != index) {
            int tmp = this->numbers[index];
            this->numbers[index] = this->numbers[smallest_node];
            this->numbers[smallest_node] = tmp;
            this->min_heapify(smallest_node);
        }
    }

public:
    KthLargest(int k, vector<int> nums) :capacity(k), numbers(vector<int>(0)) {
        for (auto& n : nums) {
            this->insert(n);
        }
    }

    int add(int val) {
        this->insert(val);
        return this->numbers.at(0);
    }
};

string judge(int except, int value) {
    return (except == value) ? "Correct" : "Incorrect: get " + std::to_string(value) + " but except " + std::to_string(except); 
}

int main() {
    int k = 3;
    vector<int> arr = {4, 5, 8, 2};

    KthLargest test1(k, arr);

    cout << judge(4, test1.add(3) ) << endl; // returns 4
    cout << judge(5, test1.add(5) ) << endl; // returns 5
    cout << judge(5, test1.add(10)) << endl; // returns 5
    cout << judge(8, test1.add(9) ) << endl; // returns 8
    cout << judge(8, test1.add(4) ) << endl; // returns 8

    KthLargest test2(1, {});

    cout << judge(-3, test2.add(-3)) << endl;
    cout << judge(-2, test2.add(-2)) << endl;
    cout << judge(-2, test2.add(-4)) << endl;
    cout << judge( 0, test2.add(0) ) << endl;
    cout << judge( 4, test2.add(4) ) << endl;
}
