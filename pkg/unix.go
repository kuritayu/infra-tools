package pkg

import (
	"bufio"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

func Ls(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	var result []string
	for _, file := range files {
		result = append(result, file.Name())
	}
	return result, err
}

func Grep(data string, keyword string) string {
	result := ""
	r := regexp.MustCompile(keyword)
	if r.MatchString(data) {
		result = data
	}
	return result
}

func Sed(data string, from string, to string) string {
	r := regexp.MustCompile(from)
	if r.MatchString(data) {
		data = r.ReplaceAllString(data, to)
	}
	return data
}

func Cat(data io.Reader) []string {
	var result []string
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}

func Cut(data string, separator string, fields ...int) string {
	var result []string
	for _, field := range fields {
		result = append(result, strings.Split(data, separator)[field-1])
	}
	return strings.Join(result, separator)
}

func Wc(data []string) int {
	return len(data)
}

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
