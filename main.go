package main

import (
	"os"
	"strings"

	"asciiartfs/fs"
)

func main() {
	fileName := fs.GetTheCorrectFile(os.Args)
	filePtr, err := os.ReadFile(fileName)
	fs.HandleError(err)
	if fileName == "thinkertoy.txt" {
		fileContents := strings.Split(string(filePtr), "\r\n")
		asciiMap := fs.MakeAsciiGraphicsMap(fileContents)
		fs.Output(os.Args[1], fileContents, asciiMap)
	} else {
		fileContents := strings.Split(string(filePtr), "\n")
		asciiMap := fs.MakeAsciiGraphicsMap(fileContents)
		fs.Output(os.Args[1], fileContents, asciiMap)
	}
}
