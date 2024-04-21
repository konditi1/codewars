package kata

func ReverseWords(str string) string {
	tempStr := ""
	revStr := ""

	for _, v := range str {
		if v == ' ' {
			// add tehe reversed string and a space if you hit a space then empty the revstr
			revStr += tempStr + " "
			tempStr = " "
		} else {
			tempStr = string(v) + tempStr
		}
	}
  return revStr + tempStr // reverse those words
}