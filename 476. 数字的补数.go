package main
func findComplement(num int) int {
	newNum := num
	ans :=1
	for newNum>0{
		ans=ans<<1
		ans++
		newNum/=2
	}
	ans=ans>>1
	return ans^num
}
func main()  {
	num:=5
	println(findComplement(num))
}
