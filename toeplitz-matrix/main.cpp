#include <iostream>
#include <vector>

// https://leetcode.com/problems/toeplitz-matrix/

using namespace std;

class Solution
{
public:
	bool isToeplitzMatrix(vector<vector<int>> &matrix)
	{
		for (size_t row_index = 1; row_index < matrix.size(); row_index++)
		{
			for (size_t col_index = 1; col_index < matrix[0].size(); col_index++)
			{
				if (matrix[row_index][col_index] != matrix[row_index - 1][col_index - 1])
					return false;
			}
		}

		return true;
	}
};

int main(int argc, char **argv)
{

	Solution s;

	vector<vector<int>> t1 = {
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2},
	};

	cout << s.isToeplitzMatrix(t1) << endl;

	vector<vector<int>> t2 = {
		{1, 2},
		{2, 2},
	};

	cout << s.isToeplitzMatrix(t2) << endl;

	vector<vector<int>> t3 = {
		{11, 74, 0, 93},
		{40, 11, 74, 7},
	};

	cout << s.isToeplitzMatrix(t2) << endl;

	return 0;
}
