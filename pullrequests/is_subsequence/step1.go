//lint:file-ignore U1000 Ignore all unused code
package issubsequence

/*
時間：10分
思っていたよりも時間がかかってしまった。
*/
func isSubsequenceStep1(s string, t string) bool {
	current := 0
	for i := 0; i < len(s); i++ {
		for {
			if current >= len(t) {
				return false
			}
			if s[i] == t[current] {
				current++
				break
			}
			current++
		}
	}
	return true
}
