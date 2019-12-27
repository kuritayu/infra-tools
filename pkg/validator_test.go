package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNotExistFile(t *testing.T) {
	path := "../../test/NotExistFile.tar"
	assert.Error(t, ValidateFile(path))
}

func TestNotReadableFile(t *testing.T) {
	path := "../../test/NotReadableFile.tar"
	assert.Error(t, ValidateFile(path))
}

func TestNotTarFile(t *testing.T) {
	path := "../../test"
	assert.Error(t, ValidateTar(path))
}

func TestGzButNotTarFile(t *testing.T) {
	path := "../../test/GzButNotTarFile.txt.gz"
	assert.Error(t, ValidateTar(path))
}

func TestTarFile(t *testing.T) {
	path := "../test/test.tar"
	actual := ValidateTar(path)
	assert.Empty(t, actual)
}

func TestTarGzFile(t *testing.T) {
	path := "../test/test2.tar.gz"
	actual := ValidateTar(path)
	assert.Empty(t, actual)
}
