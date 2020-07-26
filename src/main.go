/*
 * @Author: feilong
 * @LastEditors: feilong
 * @Description:
 * @email: feilongphone@gmail.com
 */
package main

import (
	"fmt"

	"../build/antlr4/golang/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"go.uber.org/zap"
)

func work() {
	// Setup the input
	is := antlr.NewInputStream("1 + 2 * 3")

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)

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
