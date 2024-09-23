//lint:file-ignore U1000 Ignore all unused code
package issubsequence

/*
おそらくこれが一番シンプルなのではないだろうか。
*/
func isSubsequenceStep3(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}
