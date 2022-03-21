package main
func addDigits(num int) int {
	for
	{
		var t, temp_num, sum = 0,0,0
		for i := 0; i <= num; i++ {
			temp_num *= 10
			t = num % 10
			temp_num += t
			sum += t
			if sum < 10{
				return sum
			}
		}

	}
}
func main(){
	var num int= 38
	print(addDigits(num))
}