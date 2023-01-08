
# https://leetcode.com/problems/permutation-in-string/

class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        if len(s1) > len(s2):
            return False

        s1_check = [0] * 26
        s2_check = [0] * 26

        for i, x in enumerate(s1):
            s1_check[ord(x) - ord('a')] += 1
            s2_check[ord(s2[i]) - ord('a')] += 1

        matches = 0
        for i in range(len(s1_check)):
            if s1_check[i] == s2_check[i]:
                matches += 1

        l = 0
        for r in range(len(s1), len(s2)):
            if matches == 26:
                return True

            r_index = ord(s2[r]) - ord('a')
            s2_check[r_index] += 1
            if s1_check[r_index] == s2_check[r_index]:
                matches += 1
            elif s1_check[r_index] + 1 == s2_check[r_index]:
                matches -= 1

            l_index = ord(s2[l]) - ord('a')
            s2_check[l_index] -= 1
            if s1_check[l_index] == s2_check[l_index]:
                matches += 1
            elif s1_check[l_index] - 1 == s2_check[l_index]:
                matches -= 1

            l += 1

        return matches == 26


if __name__ == "__main__":
    s = Solution()

    assert True == s.checkInclusion("ab", "eidbaooo")
    assert False == s.checkInclusion("ab", "eidboaoo")
    assert True == s.checkInclusion("adc", "dcda")
