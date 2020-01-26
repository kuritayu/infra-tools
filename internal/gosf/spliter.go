package gosf

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	r = regexp.MustCompile(`NF-(\d)`)
)

const (
	END    = 1
	TARGET = 1
	ALL    = "0"
	LAST   = "NF"
)

type Config struct {
	separator string
}

func NewConfig() *Config {
	return &Config{separator: " "}
}

func Concat(str string, config *Config, fields ...string) (string, error) {
	var result []string
	for _, field := range fields {
		target, err := SelectField(str, field, config)
		if err != nil {
			return "", err
		}
		result = append(result, target)
	}
	return strings.Join(result, config.separator), nil
}

func SelectField(str string, field string, config *Config) (string, error) {
	if field == ALL {
		return str, nil
	}

	s := strings.Split(str, config.separator)

	if field == LAST {
		return s[len(s)-END], nil
	}

	if r.MatchString(field) {
		shift, _ := strconv.Atoi(r.FindStringSubmatch(field)[TARGET])
		return s[len(s)-shift-END], nil
	}

	i, err := strconv.Atoi(field)
	if err != nil {
		return "", err
	}
	return s[i-END], nil
}
