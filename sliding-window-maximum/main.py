
# https://leetcode.com/problems/sliding-window-maximum/

from typing import List
from collections import deque


class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        ans: List[int] = []

        dq = deque()

        l = 0
        for r in range(len(nums)):
            while dq and nums[r] > nums[dq[-1]]:
                dq.pop()
            dq.append(r)

            while dq and l > dq[0]:
                dq.popleft()

            if r+1 >= k:
                ans.append(nums[dq[0]])
                l += 1

        return ans


if __name__ == "__main__":

    s = Solution()
    assert [3, 3, 5, 5, 6, 7] == s.maxSlidingWindow(
        [1, 3, -1, -3, 5, 3, 6, 7], 3)
    assert [1, -1] == s.maxSlidingWindow([1, -1], 1)
