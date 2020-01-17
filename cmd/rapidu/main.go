package main

import (
	"flag"
	"fmt"
	"github.com/kuritayu/infra-tools/internal/rapidu"
	"sync"
	"time"
)

var (
	verbose = flag.Bool("v", false, "show verbose progress messages")
)

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	sizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go rapidu.Walk(root, &n, sizes)
	}

	go func() {
		n.Wait()
		close(sizes)
	}()

	//go func() {
	//	os.Stdin.Read(make([]byte, 1))
	//	close(rapidu.Done)
	//}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, total int64

loop:
	for {
		select {
		case <-rapidu.Done:
			for range sizes {
			}
			return
		case size, ok := <-sizes:
			if !ok {
				break loop
			}
			nfiles++
			total += size
		case <-tick:
			fmt.Print(rapidu.PrintDiskUsage(nfiles, total))
		}
	}

	fmt.Print(rapidu.PrintDiskUsage(nfiles, total))
}
