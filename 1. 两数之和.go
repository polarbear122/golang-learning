package main

import "fmt"

func twoSum(arrayInput []int,target int)(a,b int){
	hashMap := map[int]int{}
	for i:=0;i<len(arrayInput);i++{
		if p,ok:=hashMap[target-arrayInput[i]];ok{
			return p,i
		}
		hashMap[arrayInput[i]]=i
	}
	return -1,-1
	//for i:=0;i<len(arrayInput);i++{
	//	hashMap[target-arrayInput[i]]=i
	//}
	//for i:=0;i<len(arrayInput);i++{
	//	if hashMap[arrayInput[i]]!=0{
	//		return i,hashMap[arrayInput[i]]
	//	}
	//}
	//return -1,-1
}
func main() {
	arr := []int{1,3,5,7}
	target:=12
	fmt.Println(twoSum(arr,target))
}
