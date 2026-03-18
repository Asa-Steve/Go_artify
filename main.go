package main

import (
	_ "embed"
	"fmt"
	"io"
	"os"
	"strings"
)

type File struct {
	content        string
	contentLines   []string
	artWordPerLine []string
}

func (f *File) PrintContentAsAscii(w io.Writer) {

	for _, wordsPerLine := range f.artWordPerLine { //["hello","there"]
		if wordsPerLine == "" {
			fmt.Println()
			continue
		}
		for line := 1; line < 9; line++ {
			for i := 0; i < len(wordsPerLine); i++ {
				char := wordsPerLine[i]

				startIdx := int((char - 32)) * 9
				fmt.Fprintf(w, string(f.contentLines[startIdx+line]))
			}
			fmt.Fprintln(w)
		}

	}

}

// reading file to memory
//
//go:embed standard.txt
var embedFileContent string

func Init(arg string, w io.Writer) error {

	// checking if the argument isn't empty and contains at least a character
	if len(arg) < 1 {
		return nil
	}

	// checking if file was loaded properly
	if len(embedFileContent) == 0 {
		fmt.Println("banner File empty or missing")
		return nil
	}

	// creating a new file instance
	var file File

	// assigning the read content to the File struct
	file.content = embedFileContent

	// splitting files into lines
	file.contentLines = strings.Split(file.content, "\n")

	// splitting art word into lines
	artWord := strings.ReplaceAll(os.Args[1], "\\n", "\n") 
	file.artWordPerLine = strings.Split(artWord, "\n") // [hello,there]

	// printing char
	file.PrintContentAsAscii(w)

	return nil

}

func main() {

	// ensuring exactly one argument is passed and nothing else
	if len(os.Args) != 2 {
		fmt.Println("Error: invalid number of argument")
		return
	}

	// initiating the program
	err := Init(os.Args[1], os.Stdout)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

}
