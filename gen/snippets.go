// +build ignore

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sourcegraph/syntaxhighlight"
)

// LinesInFile will read all lines in a file to a slice
func LinesInFile(fileName string) []string {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		result = append(result, line)
	}
	return result
}

func writeSnippetToFile(codeFile string, start int, stop int, targetFile string) {
	mainGoLines := LinesInFile("main.go")

	slice := strings.Join(mainGoLines[start-1:stop], "\n")
	highlighted, err := syntaxhighlight.AsHTML([]byte(slice))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	err = ioutil.WriteFile(targetFile, highlighted, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
}

func main() {
	writeSnippetToFile("main.go", 2, 47, "resources/snippet_i_preparations.html")
	writeSnippetToFile("main.go", 47, 64, "resources/snippet_ii_search_pod.html")
	writeSnippetToFile("main.go", 82, 105, "resources/snippet_iii_search_snapshot.html")
	writeSnippetToFile("main.go", 107, 144, "resources/snippet_iv_retrieve_metrics.html")
}
