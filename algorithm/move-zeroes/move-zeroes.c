void moveZeroes(int* nums, int numsSize) {
    int i, j, count = 0;
    for (i = 0, j = 0; i < numsSize; i++) {
        if (nums[j] != 0) {
            j++;
            continue;
        }
        if (nums[i] != 0 && nums[j] == 0) {
            nums[j] = nums[i];
            nums[i] = 0;
            j++;
            count++;
        }
    }
}
