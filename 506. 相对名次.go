package main

import (
	"fmt"
	"sort"
	"strconv"
)

//给出 N 名运动员的成绩，找出他们的相对名次并授予前三名对应的奖牌。
//前三名运动员将会被分别授予 “金牌”，“银牌” 和“ 铜牌”（"Gold Medal", "Silver Medal", "Bronze Medal"）。
//(注：分数越高的选手，排名越靠前。)

func findRelativeRanks(score []int) []string {
	scoreRank := make([]int, len(score))
	copy(scoreRank, score)

	// 降序后获取排名
	sort.Sort(sort.Reverse(sort.IntSlice(scoreRank)))
	rankMap := make(map[int]int, len(scoreRank))
	for i, v := range scoreRank {
		rankMap[v] = i+ 1
	}

	rankMedal := map[int]string{
		1: "Gold Medal",
		2: "Silver Medal",
		3: "Bronze Medal",
	}

	res := []string {}
	for _, value := range score {
		if rank, ok := rankMedal[rankMap[value]]; ok {
			res = append(res, rank)
		} else {
			res = append(res, strconv.Itoa(rankMap[value]))
		}
	}

	return res
}


func main()  {
	score:=[]int{5, 4, 3, 2, 1}
	fmt.Printf("%v",findRelativeRanks(score))
}
