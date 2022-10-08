from typing import List

# https://leetcode.com/problems/trapping-rain-water/

class Solution:
    def trap(self, height: List[int]) -> int:
        area = 0

        l, r = 0, len(height) - 1
        max_left, max_right = height[l], height[r]
        while l < r:
            if height[l] < height[r]:
                l += 1
                max_left = max(height[l], max_left)
                area += max_left - height[l]
            else:
                r -= 1
                max_right = max(height[r], max_right)
                area += max_right - height[r]

        return area


if __name__ == "__main__":
    s = Solution()
    print(s.trap([0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]))
