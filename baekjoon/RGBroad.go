package baekjoon

import (
	"fmt"
)

//백준 1149번 RGB거리 문제 풀이
func RGBroad() {
	var n int
	fmt.Scanln(&n)

	rgb := make([][]int, n)

	rgb[0] = make([]int, 3)

	fmt.Scan(&rgb[0][0])
	fmt.Scan(&rgb[0][1])
	fmt.Scan(&rgb[0][2])

	for i := 1; i < n; i++ {
		rgb[i] = make([]int, 3)

		fmt.Scan(&rgb[i][0])
		fmt.Scan(&rgb[i][1])
		fmt.Scan(&rgb[i][2])

		if rgb[i-1][2] > rgb[i-1][1] {
			rgb[i][0] += rgb[i-1][1]
		} else {
			rgb[i][0] += rgb[i-1][2]
		}

		if rgb[i-1][2] > rgb[i-1][0] {
			rgb[i][1] += rgb[i-1][0]
		} else {
			rgb[i][1] += rgb[i-1][2]
		}

		if rgb[i-1][0] > rgb[i-1][1] {
			rgb[i][2] += rgb[i-1][1]
		} else {
			rgb[i][2] += rgb[i-1][0]
		}

	}

	answer := rgb[n-1][0]
	if rgb[n-1][1] < answer {
		answer = rgb[n-1][1]
	}
	if rgb[n-1][2] < answer {
		answer = rgb[n-1][2]
	}

	println(answer)

}
