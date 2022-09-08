#include <iostream>
#include <vector>

// https://leetcode.com/problems/max-consecutive-ones/

using namespace std;

class Solution
{
public:
    int findMaxConsecutiveOnes(vector<int> &nums)
    {
        int max = 0;

        int sum = 0;
        for (auto x : nums)
        {
            if (x != 1)
            {
                max = sum > max ? sum : max;
                sum = 0;
            }
            else
            {
                sum++;
            }
        }
        max = sum > max ? sum : max;
        return max;
    }
};

int main(int argc, char **argv)
{

    Solution s;

    vector<int> v1 = {1, 1, 1, 0, 1, 1};
    cout << s.findMaxConsecutiveOnes(v1) << endl;

    return 0;
}
