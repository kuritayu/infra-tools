package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/kuritayu/infra-tools/internal/gosf"
	"log"
	"os"
	"regexp"
)

var (
	f = flag.String("f", " ", "field separator")
	r = regexp.MustCompile(`\d+|NF|NF-\d`)
)

func main() {
	// usageの定義
	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Usage: %s [-f separator] f1 f2 ...\n",
			os.Args[0])
	}
	// 引数のパース
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// 定義の作成
	config := gosf.NewConfig(*f)

	// 最終要素の取得(最終要素は、標準入力を示す"-"か、ファイル名が入っている)
	last := args[len(args)-1]

	// 最終要素が数字 or NF or NF-1の場合は、標準入力からデータを取得するため、
	// last変数は"-"にする
	// それ以外の場合は、"-" or ファイル名のはずのため、argsをフィールドだけにしたいため、
	// 最終要素をargsから除去する
	if r.MatchString(last) {
		last = "-"
	} else {
		// 最終要素を引数から除去する
		args = args[:len(args)-1]
	}

	// 対象データを読み込む
	var scanner = readData(last)

	// 対象データを引数で指定されたフィールドごとに出力する
	for scanner.Scan() {
		str := scanner.Text()
		p, err := gosf.Concat(str, config, args)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(p)
	}
}

func readData(flag string) *bufio.Scanner {
	if flag == "-" {
		return bufio.NewScanner(os.Stdin)
	} else {
		f, err := os.Open(flag)
		if err != nil {
			log.Fatalln(err)
		}
		return bufio.NewScanner(f)
	}
}
