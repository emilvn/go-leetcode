package main

// # LONGEST COMMON PREFIX
//
// Write a function to find the longest common prefix string amongst an array of strings.
//
// If there is no common prefix, return an empty string "".
//
// Example 1:
//
// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:
//
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.
func longestCommonPrefix(strs []string) string {
	control := strs[0]
	lcp := control

	for _, str := range strs {
		cp := getCommonPrefix(control, str)
		if len(cp) < len(lcp) {
			lcp = cp
		}
	}
	return lcp
}

func getCommonPrefix(s1 string, s2 string) string {
	if len(s1) == 0 || len(s2) == 0 {
		return ""
	}
	var r1 []rune
	var r2 []rune
	if len(s1) >= len(s2) {
		r1 = []rune(s1)
		r2 = []rune(s2)
	} else {
		r1 = []rune(s2)
		r2 = []rune(s1)
	}

	common := ""
	for i, r := range r2 {
		if r != r1[i] {
			return common
		}
		common += string(r)
	}
	return common
}
