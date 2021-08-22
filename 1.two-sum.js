/*
 * @lc app=leetcode id=1 lang=javascript
 *
 * [1] Two Sum
 */

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {
  let map = new Map()

  for (i = 0; i < nums.length; i++) {
    if (map.has(target - nums[i])) {
      return [map.get(target - nums[i], i)]
    } else {
      map.set(nums[i], i)
    }
  }
}
// @lc code=end
