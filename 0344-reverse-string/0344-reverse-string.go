func reverseString(s []byte) {
	var mid int
	if len(s)%2 == 0 {
		mid = (len(s) / 2) - 1
	} else {
		mid = (len(s) / 2)
	}
	_reverseString(s, mid)
}

func _reverseString(s []byte, mid int) {
	if mid == -1 {
		return
	}
	s[mid], s[len(s)-1-mid] = s[len(s)-1-mid], s[mid]
	_reverseString(s, mid-1)
}
