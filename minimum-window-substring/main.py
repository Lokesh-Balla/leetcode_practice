# https://leetcode.com/problems/minimum-window-substring/


class Solution:
    def minWindow(self, s: str, t: str) -> str:
        ans = ""
        ans_len = float("inf")

        have, need = 0, 0
        freq = dict()
        for x in t:
            freq[x] = 1 + freq.get(x, 0)
            need += 1
        window = {x: 0 for x in freq.keys()}

        l, r = 0, 0

        while r < len(s):
            if s[r] in freq.keys():
                window[s[r]] += 1
                if window[s[r]] <= freq[s[r]]:
                    have +=1

            while have == need:
                if ans_len > (r-l+1):
                    ans = s[l:r+1]
                    ans_len = r-l + 1

                if s[l] in freq.keys():
                    window[s[l]] -= 1
                    if window[s[l]] < freq[s[l]]:
                        have -= 1
                l += 1

            r += 1

        if ans is None:
            return ""

        return ans


if __name__ == "__main__":
    s = Solution()
    print(s.minWindow(s="ADOBECODEBANC", t="ABC"))
    print(s.minWindow(s="a", t="a"))
    print(s.minWindow(s="a", t="aa"))
    print(s.minWindow("bba", "ab"))
    print(s.minWindow("acbbaca","aba"))
