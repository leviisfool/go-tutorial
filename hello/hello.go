package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"bou.ke/monkey"
)

func main() {
	/*	var num, sum int
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			for _, str := range strings.Fields(sc.Text()) {
				num, _ = strconv.Atoi(str)
				sum += num
			}
			fmt.Println(sum)
			sum = 0
		}*/

	/*s := "abcbca"
	r := s[1:4]
	println(r)*/

	//s := "ab"

	/*println(reverse(123456))

	s := " 123"
	println(s[0] == ' ')
	println(s[1] - '1' == 0)*/

	monkey.Patch(fmt.Println, func(a ...interface{}) (n int, err error) {
		s := make([]interface{}, len(a))
		for i, v := range a {
			s[i] = strings.Replace(fmt.Sprint(v), "hell", "*bleep*", -1)
		}
		return fmt.Fprintln(os.Stdout, s...)
	})
	fmt.Println("what the hell?") // what the *bleep*?

}

func reverse(x int) int {
	max := strconv.Itoa(int(math.Pow(2, 31)) - 1)
	negMax := strconv.Itoa(int(math.Pow(2, 31)))

	mark := ""
	if x < 0 {
		mark = "-"
	}

	xStr := strconv.Itoa(x)
	if x < 0 {
		xStr = xStr[1:]
	}

	r := []uint8{}
	for i := len(xStr) - 1; i > 0; i-- {
		r = append(r, xStr[i])
	}
	rStr := string(r)

	if x < 0 && rStr > negMax || x > 0 && rStr > max {
		return 0
	}

	atoi, _ := strconv.Atoi(mark + rStr)
	return atoi
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	db := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		db[i] = make([]bool, len(s))
		db[i][i] = true
	}

	maxLen := 1
	begin := 0

	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				db[i][j] = false
			} else {
				if j-i < 3 {
					db[i][j] = true
				} else {
					db[i][j] = db[i+1][j-1]
				}
			}

			if db[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}
