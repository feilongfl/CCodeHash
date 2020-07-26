package main

import (
	"errors"
	"io/ioutil"
	"os"

	"CCodeHash/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

//ReadFile read file to string.
func ReadFile(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		zap.L().Info("file read failed!")
	}

	return string(content)
}

func work(path string) {
	// initial map
	FuncTableInit()

	// Setup the input
	is := antlr.NewInputStream(ReadFile(path))

	// Create the Lexer
	lexer := parser.NewCLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCParser(stream)

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&cListener{}, p.CompilationUnit())

	// print result
	PrintResult()
}

func main() {
	// initial log
	logger := zap.NewExample()
	defer logger.Sync()

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	// zap.L().Debug("zap initial finish.")

	// Parse commandline
	app := &cli.App{
		Action: func(c *cli.Context) error {
			path := "Nefertiti"
			if c.NArg() > 0 {
				path = c.Args().Get(0)
				work(path)
				return nil
			}
			return errors.New("no file!")
		},
	}

	cli.AppHelpTemplate = `Calc funchash for C code
USAGE:
	CCodeHash file
`

	app.Run(os.Args)
}
