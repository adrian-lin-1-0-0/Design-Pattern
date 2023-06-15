package main

import (
	"fmt"

	"4.2.H/pkg/analyzer"
)

func main() {
	script := `A: B C D
	B: A D E
	C: A E G K M
	D: A B K P
	E: B C J K L
	F: Z`

	a := analyzer.NewSuperRelationshipAnalyzerAdapter()
	rGraph := a.Init(script)
	fmt.Println(a.GetMutualFriends("A", "B"))
	// [D C]
	fmt.Println(rGraph.HasConnection("A", "L"))
	// true
	fmt.Println(rGraph.HasConnection("A", "F"))
	// false
}
