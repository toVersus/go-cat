package main

import (
	"testing"
)

func TestDoFileScan(t *testing.T) {
	var testFilePath = `.\content.txt`
	err := doFileScan(testFilePath)

	if err != nil {
		t.Error(err)
	}
}
