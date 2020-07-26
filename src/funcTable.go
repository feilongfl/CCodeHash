package main

import (
	"encoding/hex"
	"fmt"
	"hash/fnv"

	"go.uber.org/zap"
)

var FuncTable map[string]string

func FuncTableInit() {
	FuncTable = make(map[string]string)
}

func FuncTableAdd(sig string, body string) {
	if val, ok := FuncTable[sig]; ok {
		zap.L().Error("same func sig in same file.[" + val + "]")
	} else {
		// calc hash here
		h := fnv.New128()
		h.Write([]byte(body))
		FuncTable[sig] = hex.EncodeToString(h.Sum(nil))
	}
}

func PrintResult() {
	for k, v := range FuncTable {
		fmt.Printf("%s,%s\n", k, v)
	}
}
