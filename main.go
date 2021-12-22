package main

import (
	"os"

	"github.com/gramidt/mash-lang-for-codemash/console"
)

func main() {
	console.StartRepl(os.Stdin, os.Stdout)
}
