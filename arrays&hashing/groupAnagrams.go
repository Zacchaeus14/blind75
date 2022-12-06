package arrays_hashing

import (
	"fmt"
	"sort"
)

//49. 字母异位词分组
//
//给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
//字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。
//
//
//
//示例 1:
//
//输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
//示例 2:
//
//输入: strs = [""]
//输出: [[""]]
//
//示例 3:
//
//输入: strs = ["a"]
//输出: [["a"]]
//
//
//
//提示：
//
//1 <= strs.length <= 104
//0 <= strs[i].length <= 100
//strs[i]   仅包含小写字母

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ans := [][]string{}
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

//执行用时：16 ms, 在所有 Go 提交中击败了96.24% 的用户 O(mlogm * n) n = arr len, m = str len
//内存消耗：8.7 MB, 在所有 Go 提交中击败了41.99% 的用户 O(m*n)

func groupAnagramsV2(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		counts := make([]int, 26)
		for _, c := range str {
			counts[int(c)-97]++
		}
		mp[fmt.Sprint(counts)] = append(mp[fmt.Sprint(counts)], str)
	}
	ans := [][]string{}
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

//执行用时：92 ms, 在所有 Go 提交中击败了5.96% 的用户 O(m * n)
//内存消耗：9.6 MB, 在所有 Go 提交中击败了5.88% 的用户 O(m * n)
