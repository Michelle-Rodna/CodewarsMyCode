/*
Write a function that takes an array of numbers (integers for the tests) and a target number. It should find two different items in the array that, when added together, give the target value. The indexes of these items should then be returned in a tuple / list (depending on your language) like so: (index1, index2).

For the purposes of this kata, some tests may have multiple answers; any valid solutions will be accepted.

The input will always be valid (numbers will be an array of length 2 or greater, and all of the items will be numbers; target will always be the sum of two different items from that array).

Based on: https://leetcode.com/problems/two-sum/
*/

package kata

func TwoSum(numbers []int, target int) [2]int {
  lenS := len(numbers)
  rez := [2]int{}
  for i, j := 0, 1; i<lenS && j < lenS;j++{
    if numbers[i] + numbers[j] == target{
      rez[0], rez[1] = i, j
      return rez
    }
    if j == lenS-1{
      i++
      j=i
    }
  }
  return rez
}