package exercises

func Anagrams(s1, s2 string) bool {
	letterCount := make(map[byte]int)

	if len(s1) != len(s2) {
		return false
	}

	for pos := 0; pos < len(s1); pos++ {

		l1 := s1[pos]
		l2 := s2[pos]

		letterCount[l1] += 1
		letterCount[l2] -= 1
		if letterCount[l1] == 0 {
			delete(letterCount, l1)
		}
		if letterCount[l2] == 0 {
			delete(letterCount, l2)
		}
	}

	return len(letterCount) == 0

}
