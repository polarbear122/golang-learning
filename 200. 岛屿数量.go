package main

import "fmt"

func main() {
	nums := [][]byte{{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}
	fmt.Println(numIslands(nums))
}

func numIslands(grid [][]byte) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				ans++
				bfs(&grid, i, j)
			}
		}
	}
	return ans
}

type PII struct {
	x int
	y int
}

var dx = [...]int{-1, 0, 1, 0}
var dy = [...]int{0, 1, 0, -1}

func bfs(grid *[][]byte, x int, y int) {
	g := *grid
	que := make([]PII, 0)
	que = append(que, PII{x, y})
	g[x][y] = '0'

	for len(que) > 0 {
		t := que[0]
		que = que[1:]

		for i := 0; i < 4; i++ {
			a := t.x + dx[i]
			b := t.y + dy[i]
			if a >= 0 && a < len(g) && b >= 0 && b < len(g[0]) && g[a][b] == '1' {
				que = append(que, PII{a, b})
				g[a][b] = '0'
			}
		}
	}
}
