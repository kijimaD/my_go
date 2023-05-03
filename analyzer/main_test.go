package main

import (
	"fmt"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "a")
}

func TestEx(t *testing.T) {
	gopher := "a"
	fmt.Sprintf("%s", gopher)
}
