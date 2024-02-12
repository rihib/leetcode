package blind75

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for {
		if i >= j {
			return true
		}

		if !(('A' <= s[i] && s[i] <= 'Z') || ('a' <= s[i] && s[i] <= 'z') || ('0' <= s[i] && s[i] <= '9')) {
			i++
			continue
		}
		if !(('A' <= s[j] && s[j] <= 'Z') || ('a' <= s[j] && s[j] <= 'z') || ('0' <= s[j] && s[j] <= '9')) {
			j--
			continue
		}

		if s[i] != s[j] {
			si := s[i]
			sj := s[j]
			if 'A' <= s[i] && s[i] <= 'Z' {
				si = s[i] - 'A' + 'a'
			}
			if 'A' <= s[j] && s[j] <= 'Z' {
				sj = s[j] - 'A' + 'a'
			}

			if si != sj {
				return false
			}
		}

		i++
		j--
	}
}
