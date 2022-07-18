package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// create function to reverse string
func reverse_string(sd string) string {
	// get the length of input string
	len_sd := int(len(sd))
	// create array of string to receive the reverse string
	var new_string []string
	// This loop use receive each character
	for i := 0; i < len(sd); i++ {
		new_string = append(new_string, string(sd[len_sd-i-1]))
	}
	// Join each element in array
	output_string := strings.Join(new_string, "")
	// return
	return output_string
}

func main() {
	var whatIsIt string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)
	// sd => iangnoW:NAM:ENIL:ta:su:nioJ
	// fmt.Println(string(sd))
	whatIsIt = reverse_string(string(sd))

	// whatIsIt
	fmt.Println(whatIsIt)
}
