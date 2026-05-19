/*
You are given an array(list) strarr of strings and an integer k. Your task is to return the first longest string consisting of k consecutive strings taken in the array.
*/
package kata

import (
  "strings"
)

func LongestConsec(strarr []string, k int) string {
  lenStr := len(strarr)
  if lenStr < k || lenStr == 0 || k <= 0{
    return ""
  }
  var rez string
  for i := 0; i+k<=lenStr; i++{
    prom := strings.Join(strarr[i:i+k],"")
    if len(prom)>len(rez){
      rez = prom
    }
  }
  return rez
}