package main

import "fmt"

func startsWith(s, pre string) bool {
	if len(s) < len(pre) {
		return false
	}
	for i, r := range pre {
		if s[i] != byte(r) {
			return false
		}
	}
	return true
}

func removePrefix(s, pre string) string {
	r := []rune(s)
	return string(r[len(pre):])
}

func prefixCanSound(bab string) (bool, string) {
	l := []string{"aya", "ye", "woo", "ma"}

	for _, ss := range l {
		if startsWith(bab, ss) {
			return true, ss
		}
	}
	return false, ""
}

func canSound(bab string) bool {
	for {
		b, ss := prefixCanSound(bab)
		if !b {
			return false
		}

		bab = removePrefix(bab, ss)
		if len(bab) == 0 {
			return true
		}
	}
}

func solution(babbling []string) int {
	count := 0
	for _, bab := range babbling {
		if canSound(bab) {
			count += 1
		}
	}
	return count
}

func main() {
	sol_1 := 1
	ret_1 := solution([]string{"aya", "yee", "u", "maa", "wyeoo"})
	if sol_1 != ret_1 {
		fmt.Println("Solution 1 no pass!!!")
	}

	sol_2 := 3
	// ret_2 := solution([]string{"yemawoo"})
	ret_2 := solution([]string{"ayaye", "uuuma", "ye", "yemawoo", "ayaa"})
	if sol_2 != ret_2 {
		fmt.Println("Solution 2 no pass!!!")
	}
}
