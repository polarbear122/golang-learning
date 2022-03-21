package main
//斐波那契数，通常用 F(n) 表示，形成的序列称为 斐波那契数列 。
//该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
func fib(n int) int {
	var store [30]int
	store[0]=0
	store[1]=1
	for i:=2;i<=n;i++{
		store[i]=store[i-1]+store[i-2]
	}
	return store[n]
}
func main(){
	num:=5
	println(fib(num))
}