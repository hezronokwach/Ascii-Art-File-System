package fs

import (
	"fmt"
	"os"
	"strings"
)

// The function maps key of runes to int values. The int values represent
// the starting point of the ascii graphic.
func MakeAsciiGraphicsMap(fileContents []string) map[rune]int {
	firstAscii := 32
	asciiMap := map[rune]int{}
	for i := 0; i < len(fileContents); i++ {
		if len(fileContents[i]) == 0 || len(fileContents[i]) == 1 {
			asciiMap[rune(firstAscii)] = i + 1
			firstAscii++
		}
	}
	return asciiMap
}

// Function determining which ascii graphic file to use, based on the flag passed
func GetTheCorrectFile(args []string) string {
	if len(args) == 1 || len(args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		os.Exit(0)
	}
	if len(args) == 3 {
		if (args[2] != "standard") && (args[2] != "shadow") && (args[2] != "thinkertoy") {
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
			os.Exit(0)
		}
		switch args[2] {
		case "thinkertoy":
			return "thinkertoy.txt"
		case "shadow":
			return "shadow.txt"
		case "standard":
			return "standard.txt"
		default:
			return "standard.txt"
		}
	}
	return "standard.txt"
}

/*
Function to print the ascii graphics
Accepts string, array of filecontents and a map with ascii graphics
*/
func Output(str string, fileContents []string, asciiMap map[rune]int) {
	if len(str) == 0 {
		return
	}
	if str == "\\n" {
		fmt.Println()
		return
	}

	// check for unprintable ascii characters
	for _, char := range str {
		if ok := (char >= rune(31) && char <= rune(127)); !ok {
			fmt.Println("Unprintable ascii character(s) in input. \nUsage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
			os.Exit(0)
		}
	}

	strArr := strings.Split(str, "\\n")
	for _, v := range strArr {
		if v == "" || v == "\n" {
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range v {
				if ascii, ok := asciiMap[char]; ok {
					fmt.Print(fileContents[ascii+i])
				}
			}
			fmt.Println()
		}
	}
}

// function to output error. The function also exits the program after outputting the error
func HandleError(er error) {
	if er != nil {
		fmt.Println(er)
		os.Exit(0)
	}
}
