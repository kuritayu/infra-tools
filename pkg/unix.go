package pkg

import (
	"bufio"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

// 指定されたパスのファイルリストを返します。
func Ls(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	var result []string
	for _, file := range files {
		result = append(result, file.Name())
	}
	return result, err
}

// 入力データにキーワードが含まれていれば入力データを返します。
func Grep(data string, keyword string) string {
	result := ""
	r := regexp.MustCompile(keyword)
	if r.MatchString(data) {
		result = data
	}
	return result
}

// 入力データからfromの文字をtoに置換して返します。
func Sed(data string, from string, to string) string {
	r := regexp.MustCompile(from)
	if r.MatchString(data) {
		data = r.ReplaceAllString(data, to)
	}
	return data
}

// 入力データを読み込み、スライスで返します。
func Cat(data io.Reader) []string {
	var result []string
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}

// 入力データをセパレータで分割し、指定されたフィールドを抽出して返します。
func Cut(data string, separator string, fields ...int) string {
	var result []string
	for _, field := range fields {
		result = append(result, strings.Split(data, separator)[field-1])
	}
	return strings.Join(result, separator)
}

// スライスの要素数(行数)を返します。
func Wc(data []string) int {
	return len(data)
}

// 引数のスライスから重複要素を削除したスライスを返します。
// uniqコマンドと違い、スライスがソートされていなくても削除します。
func Uniq(data []string) []string {
	counter := make(map[string]bool)
	var result []string
	for _, value := range data {
		if !counter[value] {
			counter[value] = true
			result = append(result, value)
		}
	}
	return result
}
