/*
Given a positive number n > 1 find the prime factor decomposition of n. The result will be a string with the following form :

 "(p1**n1)(p2**n2)...(pk**nk)"
with the p(i) in increasing order and n(i) empty if n(i) is 1.

Example: n = 86240 should return "(2**5)(5)(7**2)(11)"
*/

package kata

import (
	"strconv"
	"strings"
)

func PrimeFactors(n int) string {
	f := n
	var rez strings.Builder
	for i := 2; f >= i*i; i++ { 
		if f % i == 0 {
			count := 0
			for f % i == 0 {
				count++
				f /= i
			}
			if count > 1 {
				rez.WriteByte('(')
				rez.WriteString(strconv.Itoa(i))
				rez.WriteString("**")
				rez.WriteString(strconv.Itoa(count))
				rez.WriteByte(')')
			} else {
				rez.WriteByte('(')
				rez.WriteString(strconv.Itoa(i))
				rez.WriteByte(')')
			}
		}
	}
	
	if f > 1 {
		rez.WriteByte('(')
		rez.WriteString(strconv.Itoa(f))
		rez.WriteByte(')')
	}

	return rez.String()
}