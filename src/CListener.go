package main

import (
	"CCodeHash/parser"
	"strings"
)

type cListener struct {
	*parser.BaseCListener

	//stack []int
}

func (l *cListener) EnterFunctionDefinition(c *parser.FunctionDefinitionContext) {
	// zap.L().Info("EnterFunctionDefinition!")
	var funcBody = c.GetText()
	var funcSig = funcBody[0:strings.Index(funcBody, "{")]
	FuncTableAdd(funcSig, funcBody)
	// zap.L().Info(funcSig)
	// zap.L().Info(c.GetText())
}
