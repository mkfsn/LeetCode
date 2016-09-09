class Solution {
    
    int max(int a, int b) {
        return a > b ? a : b;
    }
    
public:
    int rob(vector<int>& nums) {
        int ifRobbedPrevious = 0;
        int ifDidntRobPrevious = 0;
        
        for (int i = 0; i < nums.size(); i++) {
            int currentRobbed = ifDidntRobPrevious + nums[i];
            int currentNotRobbed = max(ifRobbedPrevious, ifDidntRobPrevious);
            
            ifRobbedPrevious = currentRobbed;
            ifDidntRobPrevious = currentNotRobbed;
        }
        
        return max(ifRobbedPrevious, ifDidntRobPrevious);
    }
};
