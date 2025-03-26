//lint:file-ignore U1000 Ignore all unused code
package generateparentheses

/*
初見で解くことができず、他の人の回答を見ながら解く形になってしまいました。
一応自分でもバックトラッキングをちゃんと理解して自力で下記のコードを書けるようにはなりました。

時間計算量と空間計算量はカタラン数になるという認識。
https://github.com/SuperHotDogCat/coding-interview/pull/7#discussion_r1577988152
https://github.com/Mike0121/LeetCode/pull/1#discussion_r1577919957
*/
func generateParenthesis_step1(n int) []string {
	var parentheses []string
	var stack []byte
	var generate func(int, int)
	generate = func(open int, closed int) {
		if open == n && closed == n {
			parentheses = append(parentheses, string(stack))
			return
		}
		if open < n {
			stack = append(stack, '(')
			generate(open+1, closed)
			stack = stack[:len(stack)-1]
		}
		if open > closed {
			stack = append(stack, ')')
			generate(open, closed+1)
			stack = stack[:len(stack)-1]
		}
	}
	generate(0, 0)
	return parentheses
}
