package main

import (
	"fmt"
)

type Vertex struct {
	lat, long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)

	m["Bell labs"] = Vertex{
		40.3, -74.3,
	}
	fmt.Println(m["Bell labs"])
}
