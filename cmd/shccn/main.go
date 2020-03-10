package main

import (
	"flag"
	"fmt"
	"github.com/kuritayu/infra-tools/pkg/shccn"
	"log"
	"sort"
)

// TODO コードメトリクスを取る
func main() {
	flag.Parse()
	args := flag.Args()

	target := args[0]

	// TODO 引数がディレクトリの場合
	sh, err := shccn.New(target)
	if err != nil {
		log.Fatalln(err)
	}

	lines := sh.GetLines()
	codes := sh.GetCodeLines()
	comments := sh.GetCommentLines()
	blanks := sh.GetBlankLines()
	functions := sh.GetFunctionLines()
	execCodes := shccn.GetCodes(sh.Lines)
	functionCodes := shccn.GetFunctions(execCodes)

	sl := make([]string, len(functionCodes))
	i := 0
	for k := range functionCodes {
		sl[i] = k
		i++
	}
	sort.Strings(sl)

	fmt.Print(shccn.BuildSummaryHeader())
	fmt.Print(shccn.BuildSummaryBody(sh.Name, lines, codes, comments, blanks, functions))
	fmt.Print(shccn.BuildFooter())

	fmt.Print(shccn.BuildFunctionHeader())
	for _, k := range sl {
		name := k
		code := len(functionCodes[k])
		ccn := shccn.CalculateCCN(functionCodes[k])
		fmt.Print(shccn.BuildFunctionBody(sh.Name, name, code, ccn))
	}
	fmt.Print(shccn.BuildFooter())
}
