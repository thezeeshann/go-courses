package main

import (
	"files/fileutils"
	"fmt"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/data/text.txt"
	c, err := fileutils.ReadFile(filePath)
	if err == nil { // nil means nothing no error
		fmt.Println(c)
		newContent := fmt.Sprintf("Hello World! %v", c)
		fileutils.WriteFile(filePath+".output.txt", newContent)
	} else {
		fmt.Printf("Error reading file %v", err)
	}

}
