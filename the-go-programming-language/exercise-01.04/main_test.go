// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.
package main

import (
	"bytes"
	"reflect"
	"testing"
)

type test struct {
	streams   []stream
	counts    map[string]int
	lineFiles map[string][]string
}

func TestCountDuplicateLines(t *testing.T) {
	tests := []test{
		{
			[]stream{
				{bytes.NewBufferString("aaa"), "file1"},
				{bytes.NewBufferString("aaa"), "file2"},
			},
			map[string]int{
				"aaa": 2,
			},
			map[string][]string{
				"aaa": {"file1", "file2"},
			},
		},
	}

	for _, test := range tests {
		counts, lineFiles := countDuplicateLines(test.streams)
		if !reflect.DeepEqual(counts, test.counts) {
			t.Errorf("got %v want %v", counts, test.counts)
		}
		if !reflect.DeepEqual(lineFiles, test.lineFiles) {
			t.Errorf("got %v want %v", lineFiles, test.lineFiles)
		}
	}
}
