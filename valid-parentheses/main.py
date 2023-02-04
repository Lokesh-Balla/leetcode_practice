# https://leetcode.com/problems/valid-parentheses/

class Solution:
    def isValid(self, s: str) -> bool:
        stack = []

        brack = {'(': ')', '{': '}', '[': ']'}

        for x in s:
            if x in ('(', '{', '['):
                stack.append(brack[x])
            if x in (')', '}', ']'):
                if (len(stack)) > 0 and x == stack[-1]:
                    stack.pop()
                else:
                    return False

        if len(stack) == 0:
            return True

        return False


if __name__ == "__main__":
    s = Solution()
    print(s.isValid("{}{}[]()"))
    print(s.isValid("()[]{}"))
    print(s.isValid("{)"))
