package utils

import (
	"strconv"
	"strings"
)
func  Cap(s []string) []string {
	for i , _ := range s {
		if strings.ContainsAny(s[i], "(cap") {
			if strings.ContainsAny(s[i],"(cap,)") {
				val, _:= strconv.Atoi(s[i+1][:len(s[i+1])])
				for x:= i- val ; x < i ; x++ {
					s = append(s[:i],s[i+1:]...)
				}
			}
		}
	}
	return  s

}