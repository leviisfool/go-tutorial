package main

import (
	"fmt"
	"github.com/namsral/flag"
)

func main() {
	var port = 3000
	flag.IntVar(&port, "port", port, "Port number")
	flag.Parse()
	fmt.Println("You seem to prefer", port)
}

func groupAnagrams(strs []string) [][]string {
	cm := map[[26]int][]string{}
	for _, str := range strs {
		key := [26]int{}
		for _, c := range str {
			key[c-'a']++
		}
		cm[key] = append(cm[key], str)
	}

	var rs = make([][]string, 0, len(cm))

	for _, v := range cm {
		rs = append(rs, v)
	}
	return rs
}
