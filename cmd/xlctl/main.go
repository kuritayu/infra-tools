package main

import (
	"flag"
	"fmt"
	"github.com/kuritayu/infra-tools/internal/xlctl"
	"log"
	"os"
)

func main() {
	// usageの定義
	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Usage: %s ls|cat filename [sheetname] \n",
			os.Args[0])

	}

	// 引数のパース
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 && len(args) != 3 {
		// 引数の数が2または3以外は終了させる
		flag.Usage()
		os.Exit(1)
	}

	// 引数取得
	operation := args[0]
	file := args[1]
	sheet := ""
	if len(args) == 3 {
		sheet = args[2]
	}

	// 構造体セット
	data, err := xlctl.NewExcel(file)
	if err != nil {
		log.Fatalln(err)
	}

	// 第一引数による分岐
	switch operation {
	case "ls":
		err := data.PrintSheetName(sheet)
		if err != nil {
			log.Fatalln(err)
		}
	case "cat":
		err := data.ConcatData(sheet)
		if err != nil {
			log.Fatalln(err)
		}
	default:
		flag.Usage()
		os.Exit(1)
	}

}
