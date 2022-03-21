package main

func main() {

}

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	bombCount := make([][]int, n)
	for i := 0; i > n; i++ {
		bombCount[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			bombCount[i][j] = (bombs[i][0]-bombs[j][0])*(bombs[i][0]-bombs[j][0]) + (bombs[i][1]-bombs[j][1])*(bombs[i][1]-bombs[j][1])
		}
	}
	return 0
}
