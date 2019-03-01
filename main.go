package main

import (
	"flag"
	"fmt"

	"github.com/y-ogura/ca-generate/generator"
)

var (
	c = flag.Bool("c", true, "create controller flag")
	u = flag.Bool("u", true, "create usecase flag")
	r = flag.Bool("r", true, "create repository flag")
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Print("\nname is required\nUsage: ca-generate [-c] [-u] [-r] name")
		return
	}
	g := generator.New([]bool{*c, *u, *r}, flag.Arg(0))
	g.Generate()

}
