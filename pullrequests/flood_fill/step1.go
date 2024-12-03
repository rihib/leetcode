//lint:file-ignore U1000 Ignore all unused code
package floodfill

/*
時間：13分
再帰を使えば良さそうというのはすぐにわかりましたが、自分にとっては実装が少し難しく感じたので、そこそこの時間がかかりました。
ただ、これぐらいの難易度の問題をちゃんと解ききり、解いた後は脳内デバッグをしてバグを見つけて取り除いた後に1発で正解することができたので、これまでのLeetCodeの練習の成果が出ているなと感じることができました。
コードについてもわりかし綺麗に書けたと思っていて、今回は特に修正すべきところを自分では見つけられませんでした。
*/
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	originalColor := image[sr][sc]
	if originalColor == color {
		return image
	}
	image[sr][sc] = color
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, direction := range directions {
		newSr, newSc := sr+direction[0], sc+direction[1]
		if 0 <= newSr && newSr < len(image) && 0 <= newSc && newSc < len(image[0]) && image[newSr][newSc] == originalColor {
			floodFill(image, newSr, newSc, color)
		}
	}
	return image
}
