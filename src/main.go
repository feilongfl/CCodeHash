package main

import (
	"fmt"
	"io/ioutil"

	"CCodeHash/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
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

func work() {
	// Setup the input
	is := antlr.NewInputStream(ReadFile("test/bt.c"))

	// Create the Lexer
	lexer := parser.NewCLexer(is)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}

func main() {
	// initial log
	logger := zap.NewExample()
	defer logger.Sync()

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	zap.L().Debug("zap initial finish.")

	// Parse commandline

	// work
	work()
}
