package repl

import (
	"bufio"
	"fmt"
	"interpreter-in-go/evaluator"
	"interpreter-in-go/lexer"
	"interpreter-in-go/object"
	"interpreter-in-go/parser"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	for {
		_, err := fmt.Fprintf(out, PROMPT)
		if err != nil {
			return
		}
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		//env := object.NewEnvironment()
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			_, err := io.WriteString(out, evaluated.Inspect())
			if err != nil {
				return
			}
			_, err = io.WriteString(out, "\n")
			if err != nil {
				return
			}
		} else {
			_, err := io.WriteString(out, program.String())
			if err != nil {
				return
			}
			_, err = io.WriteString(out, "\n")
			if err != nil {
				return
			}
		}

	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		_, err := fmt.Fprintf(out, "\t%s\n", msg)
		if err != nil {
			return
		}
	}
}
