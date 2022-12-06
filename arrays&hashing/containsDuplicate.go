package arrays_hashing

//217. 存在重复元素
//给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3,1]
//输出：true
//
//示例 2：
//
//输入：nums = [1,2,3,4]
//输出：false
//
//示例 3：
//
//输入：nums = [1,1,1,3,3,4,3,2,4,2]
//输出：true
//
//
//提示：
//
//1 <= nums.length <= 105
//-109 <= nums[i] <= 109

func containsDuplicate(nums []int) bool {
	hash := map[int]int{}
	for _, n := range nums {
		if _, ok := hash[n]; ok {
			return true
		}
		hash[n] = 1
	}
	return false
}

//执行用时：60 ms, 在所有 Go 提交中击败了89.29% 的用户
//内存消耗：9.8 MB, 在所有 Go 提交中击败了17.08% 的用户
