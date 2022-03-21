package main

import "fmt"

//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
func spiralOrder(matrix [][]int) (result []int) {
	m:=len(matrix)
	n:=len(matrix[0])
	k:=0
	top,bottom,left,right:=0,m,0,n
	result = make([]int,m*n)
	for top<bottom&&left<right{
		for i:=left;i<right;i++{
			result[k]=matrix[top][i]
			fmt.Printf("result[k]=matrix[top][i],%d,%d,%d\n", top, i, result[k])
			k++
		}
		top++
		for i:=top;i<bottom;i++{
			result[k]=matrix[i][right-1]
			fmt.Printf("result[k]=matrix[i][right-1],%d,%d,%d\n", i, right-1, result[k])
			k++
		}
		right--
		if left<right&&top<bottom{
			for i:=right-1;i>=left;i--{
				result[k]=matrix[bottom-1][i]
				fmt.Printf("result[k]=matrix[bottom-1][i],%d,%d,%d\n", bottom-1, i, result[k])
				k++
			}
			bottom--
			for i:=bottom-1;i>=top;i--{
				result[k]=matrix[i][left]
				fmt.Printf("result[k]=matrix[i][left],%d,%d,%d\n", i,left, result[k])
				k++
			}
			left++
		}
	}
	return
}
func main() {
	nums:=make([][]int,5)
	for i:=0;i<5;i++{
		nums[i]=make([]int,5)
	}
	fmt.Println(len(nums))
	k:=1
	for i:=0;i<5;i++{
		for j:=0;j<5;j++{
			nums[i][j]=k
			k++
		}
	}
	fmt.Printf("%v\n",nums)
	result:=spiralOrder(nums)
	for i:=0;i<len(result);i++{
		fmt.Println(result[i])
	}
}
