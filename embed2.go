package main

import (
	"fmt"
)

type Status int

const (
	Good Status = iota
	Tired
)

func (s Status) String() string {
	switch s {
	case Good:
		return "Good"
	case Tired:
		return "Tired..."
	default:
		return ""
	}
}

type grasshopper struct {
	status Status
}

func (g *grasshopper) HighJump() {
	fmt.Println("High Jump!")
	g.status = Tired
}

type bird struct {
	status Status
	*grasshopper
}

func main() {
	aGrasshopper := &grasshopper{
		status: Good,
	}
	aBird := &bird{
		status:      Good,
		grasshopper: aGrasshopper,
	}
	aBird.HighJump()

	fmt.Println("grasshopper is", aGrasshopper.status)
	fmt.Println("bird is", aBird.status) // 埋め込み先のオブジェクトには影響を与えないので、実行したのにGood statusになる
}
