package longestcmn

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	str := strs[0]
	var minCommonIndex int = 100
	for i := 1; i < len(strs); i++ {
		a := findPrefixCommon(str, strs[i])
		if a < minCommonIndex {
			minCommonIndex = a
		}
	}
	return str[0:minCommonIndex]
}

func findPrefixCommon(str1, str2 string) int {
	var result int
	if len(str2) < len(str1) {
		str1, str2 = str2, str1
	}
	for i := 0; i < len(str1); i++ {
		if str1[i] == str2[i] {
			result++
		} else {
			break
		}
	}
	return result
}
