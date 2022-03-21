package main
func islandPerimeter(grid [][]int) int {
	result:=0
	for i:=0;i<len(grid);i=i+1{
			for j:=0;j<len(grid[0]);j++{
			if grid[i][j]==1{
				result+=4
				if (j-1)>=0&&grid[i][j-1]==1{
					result--
				}
				if j+1<len(grid[0])&&grid[i][j+1]==1{
					result--
				}
				if i-1>=0&&grid[i-1][j]==1{
					result--
				}
				if i+1<len(grid)&&grid[i+1][j]==1{
					result--
				}
			}
		}
	}
	return result
}
func main()  {
	grid := [][]int{{0,1,0,0},{1,1,1,0},{0,1,0,0},{1,1,0,0}}
	println(islandPerimeter(grid))
}
