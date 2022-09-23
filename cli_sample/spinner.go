package main

import (
	"time"
	"github.com/briandowns/spinner"
)

func main() {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()
	time.Sleep(5 * time.Second)
	s.Stop()
}
