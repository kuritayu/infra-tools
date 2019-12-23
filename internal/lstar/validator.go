package lstar

import (
	"errors"
	"fmt"
	"github.com/aybabtme/color/brush"
	"os"
	"path/filepath"
	"regexp"
)

const (
	FileIsNotExist    = " is not exist."
	FileIsNotReadable = " is not readable."
	FileIsNotTar      = " is not tar file."
)

func Validate(path string) error {

	_, err := os.Stat(path)

	// check file is exists
	if os.IsNotExist(err) {
		return errors.New(errorMsg(FileIsNotExist, path))
	}

	// check file is readable
	file, err := os.Open(path)
	if err != nil {
		return errors.New(errorMsg(FileIsNotReadable, path))
	}
	defer file.Close()

	// check file is tar or tar.gz
	e := filepath.Ext(path)
	if e != ".tar" && e != ".gz" {
		return errors.New(errorMsg(FileIsNotTar, path))
	}

	// check file is tar.gz
	if e == ".gz" {
		if !isTarGz(path) {
			return errors.New(errorMsg(FileIsNotTar, path))
		}
	}

	return nil
}

func errorMsg(msg string, target string) string {
	return fmt.Sprintf("%s %s %s\n",
		brush.Red("[ERROR] "),
		brush.Red(target),
		brush.Red(msg))

}

func isTarGz(path string) bool {
	rep := regexp.MustCompile(`.gz$`)
	base := filepath.Base(rep.ReplaceAllString(path, ""))
	if filepath.Ext(base) == ".tar" {
		return true
	}

	return false
}
