package utils

import (
	"strconv"
	"strings"
)
func  Up(s []string) []string {
	for i , _ := range s {
		if strings.ContainsAny(s[i], "(up") {
			if strings.ContainsAny(s[i],"(up,)") {
				val, _:= strconv.Atoi(s[i+1][:len(s[i+1])])
				for x:= i- val ; x < i ; x++ {
					s = append(s[:i],s[i+1:]...)
				}
			}
		}
	}
	return  s

}