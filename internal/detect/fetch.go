package detect

import "strings"

func Fetch(message string) (string, error) {
	args := strings.Split(message, " ")
	if len(args) != 2 {
		return "Invalid number of arguments", nil
	}

	url := args[0]
	xpath := args[1]

	content := GetContent(url, xpath)

	return content, nil
}
