package main

import "fmt"

func main() {
	root := Tree{
		LeafValue: 0,
		Right: &Tree{
			LeafValue: 5,
			Right: &Tree{
				LeafValue: 6,
				Right:     nil,
				Left:      nil,
			},
			Left: nil,
		},
		Left: &Tree{
			LeafValue: 4,
			Right:     nil,
			Left:      nil,
		},
	}

	fmt.Println(root.Right.Right.LeafValue)
}
