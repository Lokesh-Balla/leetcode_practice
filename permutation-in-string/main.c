#include <stdio.h>
#include <stdbool.h>

void reset_map(int *arr, char *s1)
{
    if (arr == NULL)
        return;
    for (int i = 0; i < 26; i++)
        arr[i] = 0;

    while (*s1 != '\0')
    {
        arr[*s1 - 'a'] += 1;
        s1++;
    }
}

bool check_map(int *arr)
{
    if (arr == NULL)
        return true;
    for (int i = 0; i < 26; i++)
    {
        if (arr[i] != 0)
            return false;
    }
    return true;
}

bool checkInclusion(char *s1, char *s2)
{
    int *s1_map = (int *)malloc(sizeof(int) * 26);
    reset_map(s1_map, s1);

    int start = 0;

    for (int i = 0; s2[i] != '\0';)
    {
        if (s1_map[s2[i] - 'a'] > 0)
        {
            s1_map[s2[i] - 'a']--;
            i++;
        }
        else
        {
            if (check_map(s1_map))
                return true;
            reset_map(s1_map, s1);
            i = start + 1;
            start++;
        }
    }

    return check_map(s1_map);
}

int main()
{
    return 0;
}