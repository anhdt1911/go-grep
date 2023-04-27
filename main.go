package main

import (
	"fmt"
	"os"

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
	for _, v := range args.File {
		content, err := os.ReadFile(v)
		if err != nil {
			panic("error while reading file")
		}
		fmt.Println(string(content))
	}
}
