package rapidu

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestEntry(t *testing.T) {
	actual := entry(".")
	expected := 2
	fmt.Println(actual)
	assert.Equal(t, expected, len(actual))
}

func TestWalk(t *testing.T) {
	roots := []string{".."}
	sizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go Walk(root, &n, sizes)
	}

	go func() {
		n.Wait()
		close(sizes)
	}()

	var nfiles, total int64
	for size := range sizes {
		nfiles++
		total += size
	}

	actual := PrintDiskUsage(nfiles, total)
	expected := "29 files 0.0 MB\n"
	assert.Equal(t, expected, actual)

}
