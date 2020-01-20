package rapidu

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

const COUNTINGSEMAPHORE = 20

var (
	sema = make(chan struct{}, COUNTINGSEMAPHORE)
)

func entry(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}:
	}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil
	}
	return entries

}

func Walk(dir string, n *sync.WaitGroup, sizes chan<- int64) {
	defer n.Done()
	for _, entry := range entry(dir) {
		if entry.IsDir() {
			n.Add(1)
			sub := filepath.Join(dir, entry.Name())
			go Walk(sub, n, sizes)
		} else {
			sizes <- entry.Size()
		}
	}
}

func PrintDiskUsage(nfiles, nbytes int64) string {
	return fmt.Sprintf("%d files %.1f MB\n", nfiles, float64(nbytes)/1e6)
}
