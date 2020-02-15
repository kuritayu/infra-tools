package pkg

import (
	"io/ioutil"
	"regexp"
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

//TODO cat
