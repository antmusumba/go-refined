package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func Bin(s []string) []string {
	for i, word := range s {
		if strings.Contains(word, "(bin)") {
			if i > 0 {
				if word == "(bin)" {
					dat, _ := strconv.ParseInt(s[i-1], 2, 64)
					s[i-1] = fmt.Sprint(dat)
					s = append(s[:i], s[i+1:]...)
				}
			}
		}
	}
	return s
}
