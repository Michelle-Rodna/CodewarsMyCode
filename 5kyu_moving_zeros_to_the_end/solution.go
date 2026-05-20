/*
Write an algorithm that takes an array and moves all of the zeros to the end, preserving the order of the other elements.
*/

package kata

func MoveZeros(arr []int) []int {
  rez := make([]int, len(arr))
  for i, j := 0, 0; i < len(arr);i++{
    if arr[i] != 0{
      rez[j] = arr[i]
      j += 1
    }
  }
  return rez
}