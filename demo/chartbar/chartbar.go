// Copyright 2021, Chef.  All rights reserved.
// https://github.com/onedss/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/onedss/naza/pkg/chartbar"

	"github.com/onedss/naza/pkg/nazalog"
)

func main() {
	filename := parseFlag()
	output, err := chartbar.WithCsv(filename)
	nazalog.Assert(nil, err)
	fmt.Print(output)
}

func parseFlag() string {
	dir := flag.String("f", "", "csv filename")
	flag.Parse()
	if *dir == "" {
		flag.Usage()
		os.Exit(1)
	}
	return *dir
}
