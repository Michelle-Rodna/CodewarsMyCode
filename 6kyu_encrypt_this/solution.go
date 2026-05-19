/*
You want to create secret messages which can be deciphered by the Decipher this! kata. Here are the conditions:

Your message is a string containing space separated words.
You need to encrypt each word in the message using the following rules:
The first letter must be converted to its ASCII code.
The second letter must be switched with the last letter
Keepin' it simple: There are no special characters in the input.
*/

package kata

import (
  "strings"
  "strconv"
)

func EncryptThis(text string) string {
  if text == ""{
    return ""
  }
  mas := strings.Fields(text)
  var rez strings.Builder
  rez.Grow(len(text))
  for index, i := range mas{
    if index > 0{
      rez.WriteString(" ")
    }
    lenI := len(i)
    switch lenI{
      case 1:
        rez.WriteString(strconv.Itoa(int(i[0])))
      case 2:
        rez.WriteString(strconv.Itoa(int(i[0])) + string(i[1]))
      case 3:
        rez.WriteString(strconv.Itoa(int(i[0])) + string(i[2]) + string(i[1]))
      default:
        rez.WriteString(strconv.Itoa(int(i[0])) + string(i[lenI-1]) + string(i[2:lenI-1]) + string(i[1]))
    }
  }
  return rez.String()
}