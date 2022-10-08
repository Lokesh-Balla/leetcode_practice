from typing import List

# https://leetcode.com/problems/container-with-most-water/


class Solution:
    def maxArea(self, height: List[int]) -> int:
        l = 0
        r = len(height) - 1
        max_area = 0
        while l < r:
            area = (r-l)*min(height[r], height[l])
            if height[r] > height[l]:
                l += 1
            else:
                r -= 1
            max_area = area if area > max_area else max_area

        return max_area


if __name__ == "__main__":
    s = Solution()

    print(s.maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7]))
