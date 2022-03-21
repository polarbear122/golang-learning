package main
//给你一个字符串 s，找到 s 中最长的回文子串。
// 动态规划

// 暴力解法
func longestPalindrome(s string) string {
	left,right:=0,0
	if len(s)==1{
		return s
	}
	for i:=0;i<len(s);i++{
		for j:=len(s)-1;j>0;j--{
			isFind:=true
			for k:=i;k<=(j+i)/2;k++{
				if s[k]!=s[i+j-k]{
					isFind=false
					break
				}
			}
			if isFind&&j-i>right-left{
				right=j
				left=i
			}
		}
	}
	if right-left!=0{
		return s[left:right+1]
	}else if len(s)==0{
		return ""
	}else if len(s)==1{
		return s
	}else {
		return s[0:1]
	}
}

func main(){
	s:="aabcc"
	println(longestPalindrome(s))
}

