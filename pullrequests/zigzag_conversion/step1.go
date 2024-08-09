//lint:file-ignore U1000 Ignore all unused code
package zigzagconversion

/*
時間：34分
解法を思いつくのに時間がかかってしまいました（10分ぐらい？）。
規則性を持っており、インデックスからどの行になるかを求めることができると思ったのでこの解法になりました。
*/
func convert_step1(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var rows = make([][]rune, numRows)
	for i, r := range s {
		tmp := i % (numRows*2 - 2)
		if tmp < numRows {
			rows[tmp] = append(rows[tmp], r)
		} else {
			rows[2*numRows-tmp-2] = append(rows[2*numRows-tmp-2], r)
		}
	}
	var res []rune
	for _, row := range rows {
		res = append(res, row...)
	}
	return string(res)
}
