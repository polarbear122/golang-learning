package main
//对于一个 正整数，如果它和除了它自身以外的所有 正因子 之和相等，我们称它为 「完美数」。
//给定一个 整数 n， 如果是完美数，返回 true，否则返回 false
func checkPerfectNumber(num int) bool {
	if num <= 0 {
		return false
	}
	temp := 0
	for i := 1;i * i <= num;i++ {
		if num % i == 0 {
			temp += i
			if i * i != num {
				temp += num / i
			}
		}
	}
	return temp == num * 2
}

func main()  {
	num := 28100
	println(checkPerfectNumber(num))
}
