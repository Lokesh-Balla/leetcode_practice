# https://leetcode.com/problems/longest-repeating-character-replacement/


class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        freq = [0] * 26
        max_length = 0
        l, r = 0, 0

        for r in range(len(s)):
            index = ord(s[r])-ord('A')
            freq[index] += 1
            while (r-l+1) - max(freq) > k:
                freq[ord(s[l])-ord('A')] -= 1
                l += 1
            max_length = max(max_length, r-l+1)
        return max_length


if __name__ == "__main__":
    s = Solution()
    print(s.characterReplacement("ABAB", 2))
    print(s.characterReplacement("AABABBA", 1))
