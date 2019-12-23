package lstar

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNotExistFile(t *testing.T) {
	path := "../../test/NotExistFile.tar"
	assert.Error(t, Validate(path))
}

func TestNotReadableFile(t *testing.T) {
	path := "../../test/NotReadableFile.tar"
	assert.Error(t, Validate(path))
}

func TestNotTarFile(t *testing.T) {
	path := "../../test"
	assert.Error(t, Validate(path))
}

func TestGzButNotTarFile(t *testing.T) {
	path := "../../test/GzButNotTarFile.txt.gz"
	assert.Error(t, Validate(path))
}

func TestTarFile(t *testing.T) {
	path := "../../test/test.tar"
	actual := Validate(path)
	assert.Empty(t, actual)
}

func TestTarGzFile(t *testing.T) {
	path := "../../test/test2.tar.gz"
	actual := Validate(path)
	assert.Empty(t, actual)
}
