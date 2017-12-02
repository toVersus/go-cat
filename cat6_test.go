package main

import (
	"testing"
)

func TestPassFileScan(t *testing.T) {
	var testFilePath = `.\content.txt`
	err := doFileScan(testFilePath)

	if err != nil {
		t.Error(err)
	}
}

func TestNotPassFileScan(t *testing.T) {
	var testFilePath = `test`
	err := doFileScan(testFilePath)

	if err == nil {
		t.Error(err)
	}
}
