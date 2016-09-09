class Queue {
    stack<int> stk;
public:
    // Push element x to the back of queue.
    void push(int x) {
        stack<int> stk;
        while (!this->stk.empty()) {
            int y = this->stk.top();
            this->stk.pop();
            stk.push(y);
        }
        
        this->stk.push(x);
        while (!stk.empty()) {
            int y = stk.top();
            stk.pop();
            this->stk.push(y);
        }
    }

    // Removes the element from in front of queue.
    void pop(void) {
        this->stk.pop();
    }

    // Get the front element.
    int peek(void) {
        return this->stk.top();
    }

    // Return whether the queue is empty.
    bool empty(void) {
        return this->stk.empty();
    }
};
