/*
 * @Author: feilong
 * @LastEditors: feilong
 * @Description:
 * @email: feilongphone@gmail.com
 */
package main

import (
	"go.uber.org/zap"
)

func main() {

	// initial log
	logger := zap.NewExample()
	defer logger.Sync()

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	zap.L().Debug("zap initial finish.")
}
