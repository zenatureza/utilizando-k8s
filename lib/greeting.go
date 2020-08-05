package lib

import "fmt"

func Greeting(message string) string {
	var result = fmt.Sprintf("<b>%s</b>", message)
	return result
}
