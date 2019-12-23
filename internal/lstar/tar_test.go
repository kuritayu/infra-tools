package lstar

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupForTest() *Tar {
	path := "../../test/test.tar"
	return Setup(path)
}

func TestTarName(t *testing.T) {
	actual := setupForTest().getTarName()
	expected := "test.tar"
	assert.Equal(t, expected, actual)

}

func TestChecksum(t *testing.T) {
	actual := setupForTest().getCheckSum()
	expected := "b441b2f9a3e8a6154f60a1ef6509e9bf"
	assert.Equal(t, expected, actual)
}

func TestArchivedFileName(t *testing.T) {
	actual := setupForTest().getArchivedFileName()
	expected := []string{"./test/", "./test/test2.txt", "./test.txt"}
	assert.ElementsMatch(t, expected, actual)
}

func TestArchivedFileChecksum(t *testing.T) {
	actual := setupForTest().getArchivedChecksum()
	expected := []string{"d41d8cd98f00b204e9800998ecf8427e", "30cf3d7d133b08543cb6c8933c29dfd7", "697f3de8175d739661ce5d0f9009eec4"}
	assert.ElementsMatch(t, expected, actual)
}

func TestPrint(t *testing.T) {
	setupForTest().Print()
	Setup("../../test/test2.tar.gz").Print()
}
