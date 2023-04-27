package main

import (
	"fmt"

	"github.com/docopt/docopt-go"
)

const Usage = `
Usage: ggr <Pattern> [<File> ...]
`

type Args struct {
	Pattern string
	File    []string
}

func main() {
	opts, _ := docopt.ParseDoc(Usage)
	var args Args
	err := opts.Bind(&args)
	if err != nil {
		fmt.Println(err)

	}
}
