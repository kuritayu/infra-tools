package xlctl

import (
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
	"strings"
)

type xl struct {
	Data *xlsx.File
}

func NewExcel(path string) (*xl, error) {
	data, err := xlsx.OpenFile(path)
	if err != nil {
		return nil, err
	}
	return &xl{Data: data}, nil
}

func (x *xl) PrintSheetName(name string) error {
	// カウンタ値は、シート内容を出力したら1追加する
	// ループが回りきって0だった場合は引数のシートがないことを意味する
	counter := 0
	for _, sheet := range x.Data.Sheets {

		// シート名の指定があり、かつ対象シートだった場合は、出力してループから抜ける
		if name != "" && name == sheet.Name {
			fmt.Println(sheet.Name)
			counter += 1
			break

			// シート名の指定があり、かつ対象シートではなかった場合は、ループを継続する
		} else if name != "" && name != sheet.Name {
			continue

			// シート名の指定がない場合は出力してループを継続する
		} else {
			fmt.Println(sheet.Name)
			counter += 1
		}
	}
	return checkCounter(counter)
}

func (x *xl) ConcatData(name string) error {
	// カウンタ値は、シート内容を出力したら1追加する
	// ループが回りきって0だった場合は引数のシートがないことを意味する
	counter := 0
	for _, sheet := range x.Data.Sheets {

		// シート名の指定があり、かつ対象シートだった場合は、出力してループから抜ける
		if name != "" && name == sheet.Name {
			printCell(sheet)
			counter += 1
			break

			// シート名の指定があり、かつ対象シートではなかった場合は、ループを継続する
		} else if name != "" && name != sheet.Name {
			continue

			// シート名の指定がない場合は出力してループを継続する
		} else {
			fmt.Printf("[SheetName: %s]\n", sheet.Name)
			printCell(sheet)
			counter += 1
		}

	}
	return checkCounter(counter)
}

func printCell(sheet *xlsx.Sheet) {
	for _, row := range sheet.Rows {
		var rowData []string
		for _, cell := range row.Cells {
			rowData = append(rowData, cell.String())
		}
		fmt.Println(strings.Join(rowData, " "))
	}
}

func checkCounter(counter int) error {
	// カウンタ値が0の場合は対象シートが見つからなかった場合である
	if counter == 0 {
		return errors.New("sheet is not found")
	} else {
		return nil
	}

}
