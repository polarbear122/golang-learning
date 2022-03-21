package main

import (
	"math"
	"unicode"
)

func myAtoi(s string) int {
	result,sign:=0,1
	for i:=0;i<len(s);i++{
		if s[i]==' '{
			continue
		}
		if s[i]=='+'{
			for j:=i+1;j<len(s)&&unicode.IsNumber(rune(s[j]));j++{
				result*=10
				result+=int(s[j]-'0')
				if result>int(math.Pow(2,31)-1){
					break
				}
			}
			break
		}else if s[i]=='-' {
			sign=-1
			for j:=i+1;j<len(s)&&unicode.IsNumber(rune(s[j]));j++{
				result*=10
				result+=int(s[j]-'0')
				if result>int(math.Pow(2,31)){
					break
				}
			}
			break
		}else if unicode.IsNumber(rune(s[i])) {
			for j:=i;j<len(s)&&unicode.IsNumber(rune(s[j]));j++{
				result*=10
				result+=int(s[j]-'0')
				if result>int(math.Pow(2,31)){
					break
				}
			}
			break
		}else {
			return 0
		}
	}
	return  isOverLimit(sign*result)

}
func isOverLimit(x int)int{
	if x>int(math.Pow(2,31)-1){
		return int(math.Pow(2,31)-1)
	}else if x+int(math.Pow(2,31))<0 {
		return  -int(math.Pow(2,31))
	}else{
		return x
	}
}
func main(){
	println(myAtoi(" 9223372036854775808"))
}
