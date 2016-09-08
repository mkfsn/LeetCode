class MinStack {
    
    int min;
    vector<int> _stack;
    
public:
    /** initialize your data structure here. */
    MinStack() {
    }
    
    void push(int x) {
        if (this->_stack.size() == 0)
            this->min = x;
        else if (this->min > x)
            this->min = x;
        this->_stack.push_back(x);
    }
    
    void pop() {
        this->_stack.erase(this->_stack.end() - 1);
        auto it = this->_stack.begin();
        this->min = *it;
        for (; it != this->_stack.end(); it++) {
            if (this->min > *it)
                this->min = *it;
        }
    }
    
    int top() {
        return this->_stack.back();
    }
    
    int getMin() {
        return this->min;
    }
};

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack obj = new MinStack();
 * obj.push(x);
 * obj.pop();
 * int param_3 = obj.top();
 * int param_4 = obj.getMin();
 */
