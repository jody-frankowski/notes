// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type stream struct {
	stream io.Reader
	name   string
}

func countDuplicateLines(streams []stream) (counts map[string]int, lineFiles map[string][]string) {
	counts = make(map[string]int)
	lineFiles = make(map[string][]string)

	for _, stream := range streams {
		countLines(stream, counts, lineFiles)
	}
	return counts, lineFiles
}

func addFile(line string, filenameToAdd string, lineFiles map[string][]string) {
	found := false
	for _, filename := range lineFiles[line] {
		if filename == filenameToAdd {
			found = true
			break
		}
	}
	if !found {
		lineFiles[line] = append(lineFiles[line], filenameToAdd)
	}
}

func countLines(stream stream, counts map[string]int, lineFiles map[string][]string) {
	var line string

	input := bufio.NewScanner(stream.stream)
	for input.Scan() {
		line = input.Text()
		counts[line]++
		addFile(line, stream.name, lineFiles)
	}
	// NOTE: ignoring potential errors from input.Err()
}

func main() {
	streams := make([]stream, 0)

	files := os.Args[1:]
	if len(files) == 0 {
		streams = append(streams, stream{os.Stdin, "stdin"})
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			streams = append(streams, stream{f, file})
			defer f.Close()
		}
	}

	counts, lineFiles := countDuplicateLines(streams)

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%v: %d\t%s\n", lineFiles[line], n, line)
		}
	}
}
