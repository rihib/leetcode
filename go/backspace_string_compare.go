//lint:file-ignore U1000 Ignore all unused code
package main

func backspaceCompare(s string, t string) bool {
	i, j := len(s)-1, len(t)-1
	sCnt, tCnt := 0, 0

	for i >= 0 || j >= 0 {
		i = nextIdx(s, i, &sCnt)
		j = nextIdx(t, j, &tCnt)

		if i >= 0 && j >= 0 && s[i] != t[j] {
			return false
		}
		if (i >= 0) != (j >= 0) {
			return false
		}
		i--
		j--
	}

	return true
}

func nextIdx(str string, idx int, cnt *int) int {
	for idx >= 0 {
		if str[idx] == '#' {
			*cnt++
		} else if *cnt > 0 {
			*cnt--
		} else {
			break
		}
		idx--
	}
	return idx
}
