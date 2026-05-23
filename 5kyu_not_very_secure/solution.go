/*
In this example you have to validate if a user input string is alphanumeric. The given string is not nil/null/NULL/None, so you don't have to check that.

The string has the following conditions to be alphanumeric:

At least one character ("" is not valid)
Allowed characters are uppercase / lowercase latin letters and digits from 0 to 9
No whitespaces / underscore
*/

package kata

func alphanumeric(str string) bool {
  if str == ""{
    return false
  }
  
  run := []rune(str)
  for _, i := range run{
    switch{
      case i == ' ' || i == '_':
        return false
      case i >= '0' && i <= '9':
        continue
      case (i >= 'a' && i <= 'z') || (i >= 'A' && i <= 'Z'):
        continue
      default:
        return false
    }
  }
  return true
}