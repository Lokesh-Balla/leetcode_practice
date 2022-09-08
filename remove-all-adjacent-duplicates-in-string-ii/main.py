# https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string-ii/

class Solution:
    def removeDuplicates(self, s: str, k: int) -> str:
        st = []

        for ch in s:
            if st and st[-1][0] == ch:
                st[-1][1] += 1
            else:
                st.append([ch, 1])
            
            if st[-1][1] == k:
                st.pop()

        ans = ""
        for x in st:
            ans += (x[0] * x[1])

        return ans

if __name__ == "__main__":
    s=Solution()
    print(s.removeDuplicates("abcd",3))
    print(s.removeDuplicates("deeedbbcccbdaa", 3))
