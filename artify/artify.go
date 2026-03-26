package artify

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
			fmt.Fprintln(w)
			continue
		}
		for line := 1; line < 9; line++ {
			for i := 0; i < len(wordsPerLine); i++ {
				char := wordsPerLine[i]
				//ensuring its printable character
				if int(char) < 32 || int(char) > 126 {
					continue
				}
				startIdx := int((char - 32)) * 9
				fmt.Fprint(w, string(f.contentLines[startIdx+line]))
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
	file.content = strings.ReplaceAll(embedFileContent, "\r\n", "\n")

	// splitting files into lines
	file.contentLines = strings.Split(file.content, "\n")

	// splitting art word into lines
	artWord := strings.ReplaceAll(os.Args[1], "\\n", "\n")
	file.artWordPerLine = strings.Split(artWord, "\n") // [hello,there]

	// printing char
	file.PrintContentAsAscii(w)

	return nil

}
