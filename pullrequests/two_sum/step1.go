//lint:file-ignore U1000 Ignore all unused code
package twosum

/*
かなり前に解いたものなので詳細は忘れてしまいましたが、ナイーブにやる方法では各文字ごとに毎回リストを走査してしまうと時間計算量がO(n^2)になってしまうので、オーバーヘッドはありますが、ハッシュ化するのが良いと考えました。

また同じ要素を２回使うのを避けるために、毎回追加する前に対応する要素がないかを確認してから追加するようにしました。
*/
func twoSum_step1(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if j, ok := m[target-n]; ok {
			return []int{i, j}
		}
		m[n] = i
	}
	return nil
}
