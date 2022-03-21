package main

import "strings"

func licenseKeyFormatting(S string, K int) string {
	str := strings.ReplaceAll(S,"-","")
	for i := len(str) - K; i > 0; i -= K {
		str = str[:i] + "-" + str[i:]
	}
	return strings.ToUpper(str)
}

