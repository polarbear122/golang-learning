package main
//给定一个整数，写一个函数来判断它是否是 3 的幂次方。如果是，返回 true ；否则，返回 false 。
//整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3x
func isPowerOfThree(n int) bool {
	if n<=0{
		return false
	}
	if n ==1{
		return true
	}
	for{
		if n%3==0 {
			return isPowerOfThree(n / 3)
		}else {
			return false
		}
	}

}
func main()  {
	n := 3

	println(isPowerOfThree(n))
}
