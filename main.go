package main

import "fmt"

func findAnagrams(s string, a string) []int {
	var result []int
	m := make(map[rune]int)
	for i := 0; (i + len(a)) <= len(s); i++ {
		isAnagram := true
		for _, c := range a {
			m[c] += 1
		}
		for _, c := range s[i : i+len(a)] {
			m[c] -= 1
		}
		for _, c := range a {
			if m[c] != 0 {
				isAnagram = false
			}
			m[c] = 0
		}
		if isAnagram {
			result = append(result, i)
		}
	}
	return result
}

func anagramTest() {
	strs := []string{"cbaebabacd", "cbaebaaacd", "abaebabaab", "abab", ""}
	a := []string{"abc", "aaa", "aab", "ab", ""}

	for i, s := range strs {
		fmt.Println("Anagrams @", findAnagrams(s, a[i]))
	}
}

func permutation(r []rune, i int, cnt *int, result map[string]bool) {
	if i >= len(r) {
		if _, found := result[string(r)]; found {
			panic(fmt.Sprintf("Duplicate permutation %s", string(r)))
		}
		*cnt++
		return
	} else {
		for j := 0; j <= i; j++ {
			r[j], r[i] = r[i], r[j]
			permutation(r, i+1, cnt, result)
			r[j], r[i] = r[i], r[j]
		}
	}
}

func fact(x int) int {
	if x <= 1 {
		return 1
	}
	return x * fact(x-1)
}

func permutationTest() {
	strs := []string{"ABCDEFGHIJKLMNOPQRSTUV", "A", ""}
	for _, s := range strs {
		cnt := 0
		result := make(map[string]bool)
		permutation([]rune(s), 1, &cnt, result)
		expect := fact(len(s))
		if cnt != expect {
			panic(fmt.Sprintf("cnt = %d, expected = %d", cnt, expect))
		}
	}
}

func main() {
	//anagramTest()
	permutationTest()
}
