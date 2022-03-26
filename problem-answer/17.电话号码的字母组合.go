package main

func main() {

}

func numCombine(str string) []string {
	strMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var result []string
	for _, v := range strMap {
		for i := 0; i < len(result); i++ {
			for j := 0; j < len(strMap[v]); j++ {
				result = append(result)
			}
		}
	}
	return []string{""}
}
