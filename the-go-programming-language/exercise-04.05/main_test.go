package main

import (
	"reflect"
	"testing"
)

func TestEliminateAdjacentDuplicates(t *testing.T) {
	input := []string{"", "1", "1", "1", "2", "2", "3", "4"}
	want := []string{"", "1", "2", "3", "4"}

	eliminateAdjacentDuplicates(&input)
	if !reflect.DeepEqual(input, want) {
		t.Errorf("have %v want %v", input, want)
	}
}
