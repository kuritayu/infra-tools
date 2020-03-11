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

	// 指定されたパスがディレクトリかどうか判定する
	isDir := false
	if stat.IsDir() {
		isDir = true
	}

	// ファイルリストを作成する
	files, err := createFileList(path, isDir)
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		var sh *shccn.FileContents
		switch isDir {
		case true:
			sh, err = shccn.New(filepath.Join(path, file.Name()))
		case false:
			sh, err = shccn.New(filepath.Join(filepath.Dir(path), file.Name()))
		}
		if err != nil {
			continue
		}

		if len(sh.Lines) != 0 && isShell(sh.Lines[0]) {
			targets = append(targets, sh)
		}
	}

	if len(targets) > 0 {
		printSummary(targets)
		printDetail(targets)
	}
}

// ファイルリストを作成する
func createFileList(path string, isDir bool) (files []os.FileInfo, err error) {
	if isDir {
		files, err = ioutil.ReadDir(path)
		if err != nil {
			return nil, err
		}
	} else {
		s, err := os.Stat(path)
		if err != nil {
			return nil, err
		}
		files = append(files, s)
	}
	return files, nil
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

// サマリ情報を出力する
func printSummary(targets []*shccn.FileContents) {
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
}

// 詳細情報を出力する
func printDetail(targets []*shccn.FileContents) {
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
