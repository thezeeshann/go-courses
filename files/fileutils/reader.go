package fileutils

import "os"

func ReadFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)

	if err != nil {
		// 1 return ""
		// panic(("ahhhhhhhhhhhh"))
		return "", err
	} else {
		return string(content), nil
	}
}
