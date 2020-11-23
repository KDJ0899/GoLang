package baekjoon

import (
	"fmt"
)

//Bulk 백준 7568번 덩치 문제 풀이.
func Bulk() {
	var n int

	fmt.Scanln(&n)

	people := make([][]int, n)
	ranks := make([]int, n)

	for i := 0; i < n; i++ {
		people[i] = make([]int, 2)

		fmt.Scan(&people[i][0])
		fmt.Scan(&people[i][1])
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if people[i][0] > people[j][0] && people[i][1] > people[j][1] {
				ranks[j]++
			} else if people[i][0] < people[j][0] && people[i][1] < people[j][1] {
				ranks[i]++
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Print(ranks[i] + 1)
		fmt.Print(" ")
	}

}
