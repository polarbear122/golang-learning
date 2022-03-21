package main
func convert(s string, numRows int) string {
	circleNum:=(numRows-1)*2//代表有几个循环，也是循环开头的数组下标
	if circleNum==0||len(s)<numRows{
		return s
	}
	resultString:=[]byte{}
	for i:=0;i<numRows;i++{
		if i==0{
			for j:=0;j<=len(s)/circleNum;j++{
				if circleNum*j<len(s){
					resultString=append(resultString,s[circleNum*j])
				}
			}
		}else if i<numRows-1{
			resultString=append(resultString,s[i])
			for j:=1;j<=len(s)/circleNum+1;j++ {
				println( circleNum*j-(i))
				println( circleNum*j+(i))
				if circleNum*j-(i)>0&&circleNum*j-(i)<len(s){
					resultString = append(resultString, s[circleNum*j-(i)])
				}
				if circleNum*j+(i)>0&&circleNum*j+(i)<len(s){
					resultString = append(resultString, s[circleNum*j+(i)])
				}
			}
		}else{
			for j:=0;j<=len(s)/circleNum;j++{
				if circleNum*j+numRows-1>0&&circleNum*j+numRows-1<len(s){
					resultString=append(resultString,s[circleNum*j+numRows-1])
				}
			}
		}
	}
	return string(resultString)
}
func main(){
	s := "pabcd"
	numRows := 4
	println(convert(s,numRows))
	result:="pabdc"
	if convert(s,numRows)==result{
		println("true")
	}
}
