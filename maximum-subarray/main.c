#include<stdio.h>

// https://leetcode.com/problems/maximum-subarray/

int maxSubArray(int* nums, int numsSize){

    int curr_sum = 0;
    int max = nums[0];
    
    for(int i=0;i<numsSize;i++) {
        if(curr_sum < 0) {
            curr_sum = 0;
        }
        curr_sum += nums[i];
        max = curr_sum > max?curr_sum:max;
    }
    
        
        
    return max;
    
}

int main() {
	return 0;
}
