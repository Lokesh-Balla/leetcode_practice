# https://leetcode.com/problems/3sum/

from typing import List

class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()

        ans = []

        for i, a in enumerate(nums):
            if i - 1 >= 0 and nums[i-1] == a:
                continue
            j = i+1
            k = len(nums)-1
            while j < k and j < len(nums) and k < len(nums):
                if a + nums[j] + nums[k] == 0:
                    ans.append([nums[i], nums[j], nums[k]])
                    j+=1
                    while j < len(nums) and nums[j] == nums[j-1] and j < k:
                        j += 1
                elif a + nums[j] + nums[k] > 0:
                    k-=1
                elif a + nums[j] + nums[k] < 0:
                    j+=1
        return ans

if __name__ == "__main__":
    s = Solution()
    print(s.threeSum([-1,0,1,2,-1,-4]))
    print(s.threeSum([0,1,1]))
    print(s.threeSum([0,0,0]))