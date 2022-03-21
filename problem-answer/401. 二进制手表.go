package main

import (
	"fmt"
	"math/bits"
)

//二进制手表顶部有 4 个 LED 代表 小时（0-11），底部的 6 个 LED 代表 分钟（0-59）。
//每个 LED 代表一个 0 或 1，最低位在右侧。

func readBinaryWatch(turnedOn int) (ans []string) {
	for h := uint8(0); h < 12; h++ {
		for m := uint8(0); m < 60; m++ {
			if bits.OnesCount8(h)+bits.OnesCount8(m) == turnedOn {
				ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return
}

func main(){
	turnedOn := 1
	println(readBinaryWatch(turnedOn))
}