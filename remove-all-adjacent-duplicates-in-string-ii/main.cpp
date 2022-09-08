#include <iostream>
#include <stack>
#include <string>

// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string-ii/

using namespace std;

class Solution
{
public:
    string removeDuplicates(string s, int k)
    {
        stack<pair<char, int>> st;
        for (auto &ch : s)
        {
            if (st.empty())
            {
                st.push(make_pair(ch, 1));
                continue;
            }
            pair<char, int> val = st.top();
            if (val.first == ch)
            {
                val.second += 1;
                st.pop();
                if (val.second < k)
                    st.push(val);
            }
            else
            {
                st.push(make_pair(ch, 1));
            }
        }

        string ans = "";
        while (!st.empty())
        {
            pair<char, int> val = st.top();
            for (int i = 0; i < val.second; i++)
                ans = val.first + ans;
            st.pop();
        }
        return ans;
    }
};

int main(int argc, char **argv)
{
    Solution s;

    cout << s.removeDuplicates("abcd", 3) << endl;
    cout << s.removeDuplicates("deeedbbcccbdaa", 3) << endl;

    return 0;
}