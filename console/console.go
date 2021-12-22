package console

import (
	"bufio"
	"fmt"
	"io"

	"github.com/gramidt/mash-lang-for-codemash/parser"
	"github.com/gramidt/mash-lang-for-codemash/scanner"
	"github.com/gramidt/mash-lang-for-codemash/types"
)

const (
	welcomeMessage = "Mashlang 1.0.0\n"
	prompt         = ">>> "
)

func StartRepl(in io.Reader, out io.Writer) {
	bufScanner := bufio.NewScanner(in)
	env := types.NewEnv()

	fmt.Fprint(out, welcomeMessage)

	for {
		fmt.Fprint(out, prompt)

		inputScanned := bufScanner.Scan()
		if !inputScanned {
			return
		}

		line := bufScanner.Text()
		scanner := scanner.NewScanner(line)
		parser := parser.NewParser(scanner)

		ast := parser.Parse()

		if len(parser.Errors()) != 0 {
			_, _ = io.WriteString(out, " parser errors:\n")
			for _, msg := range parser.Errors() {
				_, _ = io.WriteString(out, "\t"+msg+"\n")
			}
		}

		evaluated := types.Eval(ast, env)
		if evaluated != nil {
			_, _ = io.WriteString(out, evaluated.Inspect())
			_, _ = io.WriteString(out, "\n")
		}
	}
}
