#include <iostream>
#include <vector>
#include <queue>
#include <unordered_map>

using namespace std;

class Solution
{
public:
    vector<int> topKFrequent(vector<int> &nums, int k)
    {
        priority_queue<pair<int, int>> pq;

        unordered_map<int, int> m;
        for (auto &it : nums)
        {
            m[it] += 1;
        }

        for (auto x : m)
        {
            pq.push(make_pair(x.second, x.first));
        }

        vector<int> ans;

        while (k > 0)
        {
            ans.push_back(pq.top().second);
            pq.pop();
            k--;
        }

        return ans;
    }
};

int main(int argc, char **argv)
{
    return 0;
}