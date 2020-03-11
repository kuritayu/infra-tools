package main

import (
	"flag"
	"fmt"
	"github.com/kuritayu/infra-tools/pkg/shccn"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
)

// TODO コードメトリクスを取る
func main() {
	flag.Parse()
	args := flag.Args()
	path := args[0]
	var targets []*shccn.FileContents

	//パスが存在しない場合はエラー
	stat, err := os.Stat(path)
	if err != nil {
		log.Fatalln(nil)
	}
	isDir := false
	if stat.IsDir() {
		isDir = true
	}

	// TODO 冗長なコード
	if isDir {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatalln(err)
		}
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			sh, err := shccn.New(filepath.Join(path, file.Name()))
			if err != nil {
				continue
			}

			//TODO 空ファイルのときにNPE
			if isShell(sh.Lines[0]) {
				targets = append(targets, sh)
			}
		}
	} else {
		sh, err := shccn.New(path)
		if err != nil {
			log.Fatalln(err)
		}
		if isShell(sh.Lines[0]) {
			targets = append(targets, sh)
		}
	}

	// サマリ部分のループ
	fmt.Print(shccn.BuildSummaryHeader())
	for _, target := range targets {
		lines := target.GetLines()
		codes := target.GetCodeLines()
		comments := target.GetCommentLines()
		blanks := target.GetBlankLines()
		functions := target.GetFunctionLines()
		fmt.Print(shccn.BuildSummaryBody(target.Name, lines, codes, comments, blanks, functions))
	}
	fmt.Print(shccn.BuildFooter())

	// 関数部分のループ
	fmt.Print(shccn.BuildFunctionHeader())
	for _, target := range targets {
		execCodes := shccn.GetCodes(target.Lines)
		functionCodes := shccn.GetFunctions(execCodes)
		sortedKeys := sortKey(functionCodes)

		for _, k := range sortedKeys {
			name := k
			code := len(functionCodes[k])
			ccn := shccn.CalculateCCN(functionCodes[k])
			fmt.Print(shccn.BuildFunctionBody(target.Name, name, code, ccn))
		}
	}
	fmt.Print(shccn.BuildFooter())
}

// キーをソートする
func sortKey(codes map[string][]string) (result []string) {
	result = make([]string, len(codes))
	i := 0
	for k := range codes {
		result[i] = k
		i++
	}
	sort.Strings(result)
	return result
}

// 対象パスがbashかどうか判定する
func isShell(code string) bool {
	exp := regexp.MustCompile(`#!.*/sh|#!.*/bash|#!.* bash`)
	if exp.MatchString(code) {
		return true
	}
	return false
}
