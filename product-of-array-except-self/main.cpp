#include <iostream>
#include <vector>

using namespace std;

class Solution
{
public:
    vector<int> productExceptSelf(vector<int> &nums)
    {

        int n = nums.size();

        vector<int> ans(n, 0);

        int l = 1;
        int r = 1;
        for (int i = 0; i < n; i++)
        {
            ans[i] = l;
            l *= nums[i];
        }

        for (int i = n - 1; i >= 0; i--)
        {
            ans[i] *= r;
            r *= nums[i];
        }

        return ans;
    }
};

int main(int argc, char **argv) {
    
    Solution s;

    vector<int> a = {1,2,3,4};
    vector<int> b = s.productExceptSelf(a);

    for(auto x: b)
        cout << x << endl;
    
    return 0;
}