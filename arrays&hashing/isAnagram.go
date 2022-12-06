package arrays_hashing

//242. 有效的字母异位词
//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
//注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
//
//
//
//示例 1:
//
//输入: s = "anagram", t = "nagaram"
//输出: true
//
//示例 2:
//
//输入: s = "rat", t = "car"
//输出: false
//
//
//
//提示:
//
//1 <= s.length, t.length <= 5 * 104
//s 和 t 仅包含小写字母
//
//
//
//进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/valid-anagram
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func isAnagram(s string, t string) bool {
	counts := map[rune]int{}
	for _, c := range s {
		counts[c]++
	}
	for _, c := range t {
		counts[c]--
	}
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}
	return true
}

//执行用时：4 ms, 在所有 Go 提交中击败了69.48% 的用户
//内存消耗：2.6 MB, 在所有 Go 提交中击败了34.67% 的用户
