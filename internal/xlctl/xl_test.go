package xlctl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewExcel(t *testing.T) {
	excel := "/Users/kuritayu/go/src/github.com/kuritayu/infra-tools/test/testdata.xlsx"
	_, err := NewExcel(excel)
	assert.Nil(t, err)
}

func ExampleXl_PrintSheetName() {
	excel := "/Users/kuritayu/go/src/github.com/kuritayu/infra-tools/test/testdata.xlsx"
	data, _ := NewExcel(excel)
	_ = data.PrintSheetName("")
	// Output:
	// Sheet1
	// Sheet2
}

func ExampleXl_PrintSheetNameNotFound() {
	excel := "/Users/kuritayu/go/src/github.com/kuritayu/infra-tools/test/testdata.xlsx"
	data, _ := NewExcel(excel)
	_ = data.PrintSheetName("Sheet3")
	// Output:
}

func ExampleXl_PrintSheetNameFound() {
	excel := "/Users/kuritayu/go/src/github.com/kuritayu/infra-tools/test/testdata.xlsx"
	data, _ := NewExcel(excel)
	_ = data.PrintSheetName("Sheet1")
	// Output:
	// Sheet1
}

func ExampleXl_ConcatData() {
	excel := "/Users/kuritayu/go/src/github.com/kuritayu/infra-tools/test/testdata.xlsx"
	data, _ := NewExcel(excel)
	_ = data.ConcatData("Sheet1")
	// Output:
	// No Name
	// 1 Yuusuke
	// 2 Akiko
	// 3 Daiki
}

func ExampleXl_ConcatDataNotFound() {
	excel := "/Users/kuritayu/go/src/github.com/kuritayu/infra-tools/test/testdata.xlsx"
	data, _ := NewExcel(excel)
	_ = data.ConcatData("Sheet3")
	// Output:
}

func ExampleXl_ConcatDataAllSheet() {
	excel := "/Users/kuritayu/go/src/github.com/kuritayu/infra-tools/test/testdata.xlsx"
	data, _ := NewExcel(excel)
	_ = data.ConcatData("")
	// Output:
	// [SheetName: Sheet1]
	// No Name
	// 1 Yuusuke
	// 2 Akiko
	// 3 Daiki
	// [SheetName: Sheet2]
}
