package validparent

func IsValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	var arrStr string
	for _, str := range s {
		if isOpen(string(str)) {
			arrStr += string(str)
		} else if len(arrStr) > 0 && isMatch(arrStr[len(arrStr)-1:len(arrStr)], string(str)) {
			arrStr = arrStr[:len(arrStr)-1]
		} else {
			return false
		}
	}
	return len(arrStr) == 0
}

func isOpen(s string) bool {
	l := make(map[string]bool, 3)
	l["("] = true
	l["{"] = true
	l["["] = true
	return l[s]
}

func isMatch(open, close string) bool {
	switch close {
	case "}":
		return open == "{"
	case "]":
		return open == "["
	case ")":
		return open == "("
	}
	return false
}
