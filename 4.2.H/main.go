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
	E: B C J K L`

	a := analyzer.NewSuperRelationshipAnalyzerAdapter()
	a.Init(script)
	fmt.Println(a.GetMutualFriends("A", "B"))
}
