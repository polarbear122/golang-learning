package main
//给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。
func countBits(n int) []int {
	result := make([]int, n+1)
	for i:=0;i<n+1;i++{
		var num,t,sum=0,i,0

		for t > 0 {
			num = t%2
			t/=2
			if num == 1{
				sum++
			}
		}
		result[i] = sum
	}
	return result
}
func main()  {
	n := 5
	i := countBits(n)
	for j:=0;j<len(i);j++{
		println(i[j])
	}
}
