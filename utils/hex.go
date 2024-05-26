package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func Hex(s []string) []string {
	for i, word := range s {
		if strings.Contains(word, "(hex)") {
			if i > 0 {
				if word == "(hex)" {
					dat, _ := strconv.ParseInt(s[i-1], 16, 64)
					s[i-1] = fmt.Sprint(dat)
					s = append(s[:i], s[i+1:]...)
				}
			}
		}
	}
	return s
}
