package main

import (
	"context"
	"os"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/squash/ci/internal/extconfig"
)

// This program does three things:
// 1. reads and hashes these files
// 2. modifies the version field of this file
// 3. produces this file
func main() {

	ctx := context.TODO()
	if len(os.Args) != 2 {
		contextutils.LoggerFrom(ctx).Fatal("Must pass a single argument ( version )")
	}
	version := os.Args[1]

	extconfig.MustPrepareReleaseJsPackage(ctx, version)
	extconfig.MustPrepareReleaseConfigFile(ctx, version)

}
