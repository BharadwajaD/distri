package main

import (
	"fmt"
	"github.com/bharadwajaD/distri/pkg/system"
)

func add(a, b int) {
	fmt.Println(a + b)
}

func main() {

	pcount := 4
	system := system.NewSystems(pcount)

	for i := 0; i < pcount; i++ {
		system.Processes[i].Exec(func() { add(i, i+1) })
	}
}
