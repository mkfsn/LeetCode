class Stack {
    queue<int> qu;
public:
    // Push element x onto stack.
    void push(int x) {
        queue<int> qu;
        qu.push(x);
        while (!this->qu.empty()) {
            int y = this->qu.front();
            this->qu.pop();
            qu.push(y);
        }
        this->qu = qu;
    }

    // Removes the element on top of the stack.
    void pop() {
        this->qu.pop();
    }

    // Get the top element.
    int top() {
        return this->qu.front();
    }

    // Return whether the stack is empty.
    bool empty() {
        return this->qu.empty();
    }
};
