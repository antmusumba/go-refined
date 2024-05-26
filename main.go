package main

import (
	"fmt"
	"os"
	"strings"

	"refined/utils"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please Provide Enough arguments:")
		return
	}
	myArgs := os.Args[1:]
	file, err := os.ReadFile(myArgs[0])
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}
	mySlice := strings.Split(string(file), " ")
	dat := utils.Bin(mySlice)
	dat1 := utils.Hex(dat)
	final1 := strings.Join(dat1, " ")
	final2 := []byte(final1)
	os.WriteFile(myArgs[1], final2, 0o644)
}
