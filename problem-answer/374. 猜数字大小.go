package main/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guess(num int) int{
	nums := 100
	if num < nums{
		return 1
	}else if num>nums {
		return -1
	}else {
		return 0
	}
}
func guessNumber(n int) int {
	left,right := 1,n
	for left<right{
		mid := (left+right)/2
		if guess(mid) == 0{
			return mid
		}else if guess(mid) == 1{
			left = mid+1
		}else{
			right = mid-1
		}
	}
	return left
}
func main(){
	println(guessNumber(100))
}