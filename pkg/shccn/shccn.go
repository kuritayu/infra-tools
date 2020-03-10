package shccn

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// ファイルの内容を保持する構造体
type FileContents struct {
	Name  string
	Lines []string
}

// サマリレポートのセパレータ文字列
var separator = strings.Repeat("-", 80)

// ファイルの行数を返却する
func (fc *FileContents) GetLines() int {
	return len(fc.Lines)
}

// ファイルの空行数を返却する
func (fc *FileContents) GetBlankLines() (count int) {
	for _, v := range fc.Lines {
		if isBlankLine(v) {
			count++
		}
	}
	return count
}

// ファイルのコメント行数を返却する
func (fc *FileContents) GetCommentLines() (count int) {
	for i, v := range fc.Lines {
		// shebangはCodeとみなし、skip
		if i == 0 {
			continue
		}
		if isCommentLine(v) {
			count++
		}
	}
	return count
}

// ファイルのコード行数を返却する
func (fc *FileContents) GetCodeLines() int {
	total := fc.GetLines()
	blanks := fc.GetBlankLines()
	comments := fc.GetCommentLines()
	return total - blanks - comments
}

// ファイルの関数の数を返却する
func (fc *FileContents) GetFunctionLines() (count int) {
	for _, v := range fc.Lines {
		if isFunctionLine(v, true) {
			count++
		}
	}
	return count
}

// ファイルの内容を構造体に格納する
func New(path string) (*FileContents, error) {
	s, err := toSlice(path)
	if err != nil {
		return nil, err
	}

	return &FileContents{
		Name:  filepath.Base(path),
		Lines: s,
	}, nil
}

// ファイルの内容をスライスに格納する
func toSlice(path string) (result []string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}

// ファイルの内容からコメント行、空白行を削除したスライスを作成する
func GetCodes(code []string) (result []string) {
	exp := regexp.MustCompile(`\\s*$`)
	flag := false
	for i, v := range code {
		if isCommentLine(v) {
			continue
		}
		if isBlankLine(v) {
			continue
		}
		if flag {
			v = code[i-1] + v
			flag = false
		}
		if exp.MatchString(v) {
			flag = true
			continue
		}
		result = append(result, v)
	}
	return result
}

// ファイルの内容から関数行毎の辞書を作成する
func GetFunctions(code []string) map[string][]string {
	result := make(map[string][]string)
	flag := false
	var funcname string
	for _, v := range code {
		// 関数フラグが立っている場合、要素は無条件に追加してよい
		// ただし}が含まれていれば、関数フラグをfalseにしておく
		if flag {
			if isFunctionLine(v, false) {
				flag = false
				funcname = ""
				continue
			}
			result[funcname] = append(result[funcname], v)
			continue
		}

		// 関数フラグが立っていない場合
		// {が含まれていれば、要素を追加してよい
		if isFunctionLine(v, true) {
			funcname = GetFunctionName(v)
			flag = true
			continue
		}

		// {含まれていなければ、メイン処理扱い
		funcname = "BARE_CODE"
		result[funcname] = append(result[funcname], v)
	}

	return result
}

// 関数名を取得する
func GetFunctionName(line string) string {
	r := strings.NewReplacer("function", "", "{", "", "(", "", ")", "", " ", "")
	return r.Replace(line)
}

// サマリのヘッダ部を組み立てる
func BuildSummaryHeader() string {
	return fmt.Sprintf("%s\n"+
		"%-20s %10s %10s %10s %10s %10s\n"+
		"%s\n",
		separator, "Name", "Lines", "Code", "Comments", "Blanks", "Functions", separator)
}

// サマリのボディ部を組み立てる
func BuildSummaryBody(name string, lines, code, comments, blanks, functions int) string {
	return fmt.Sprintf("%-20s %10d %10d %10d %10d %10d\n",
		name, lines, code, comments, blanks, functions)
}

// フッタ部を組み立てる
func BuildFooter() string {
	return fmt.Sprintf("%s\n", separator)
}

// サイクロマティック複雑度を算出する
func CalculateCCN(code []string) (result int) {
	result = 1
	ccnExp := regexp.MustCompile(`if|while|for|;;`)
	conditionExp := regexp.MustCompile(`&&|\|\|`)
	for _, v := range code {
		tmp := strings.Replace(v, `"`, `'`, -1)
		if strings.Contains(tmp, `'`) {
			v = removeQuote(tmp)
		}

		if ccnExp.MatchString(v) {
			result++
		}

		for _, element := range strings.Split(v, " ") {
			if conditionExp.MatchString(element) {
				result++
			}
		}
	}
	return result
}

// 関数のヘッダ部を組み立てる
func BuildFunctionHeader() string {
	return fmt.Sprintf("%s\n"+
		"%-30s %20s %20s\n"+
		"%s\n",
		separator, "Name", "Code", "CCN", separator)
}

// 関数のボディ部を組み立てる
func BuildFunctionBody(script, name string, code, ccn int) string {
	return fmt.Sprintf("%-30s %20d %20d\n",
		name+"@"+script, code, ccn)
}

// 空行かどうか判定する
func isBlankLine(line string) bool {
	tmp := strings.Replace(line, " ", "", -1)
	if len(tmp) == 0 {
		return true
	}
	return false
}

// コメント行かどうか判定する
func isCommentLine(line string) bool {
	commentExp := regexp.MustCompile(`^\s*#`)
	if commentExp.MatchString(line) {
		return true
	}
	return false
}

// 関数行かどうか判定する
func isFunctionLine(line string, flag bool) bool {
	tmp := strings.Replace(line, `"`, `'`, -1)
	if strings.Contains(tmp, `'`) {
		line = removeQuote(tmp)
	}

	// flagがtrueの場合は関数が開始かどうか、falseの場合は終了かどうか
	exp := `}`
	if flag {
		exp = `.*\(\s*\)\s*{`
	}

	functionExp := regexp.MustCompile(exp)
	if functionExp.MatchString(line) {
		return true
	}
	return false
}

// クオートに挟まれた文字列を除去する
func removeQuote(line string) (result string) {
	splited := strings.Split(line, "")
	var tmp []string
	flag := false
	for _, v := range splited {
		if flag {
			if strings.Contains(v, `'`) {
				flag = false
				continue
			}
			continue
		}
		if strings.Contains(v, `'`) {
			flag = true
			continue
		}
		tmp = append(tmp, v)
	}
	return strings.Join(tmp, "")
}
