package main

func topKFrequent(nums []int, k int) []int {

	countMap := map[int]int{}
	for _, x := range nums {
		countMap[x] += 1
	}

	countArr := make([][]int, len(nums)+1)
	for k, v := range countMap {
		if len(countArr[v]) == 0 {
			countArr[v] = []int{k}
		} else {
			countArr[v] = append(countArr[v], k)
		}
	}

	ans := []int{}
	for i := len(nums); i >= 0; i-- {
		if len(countArr[i]) >0{
			for _, x := range countArr[i] {
				if k == 0 {
					return ans
				}
				if k > 0 {
					ans = append(ans, x)
				}
				k--
			}
		}
	}

	return ans
}

func main() {

}
