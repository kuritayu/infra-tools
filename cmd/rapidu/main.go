package main

import (
	"flag"
	"fmt"
	"github.com/kuritayu/infra-tools/internal/rapidu"
	"os"
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

	if len(roots) > 1 {
		printUsage()
	}

	sizes := make(chan int64)
	var n sync.WaitGroup
	root := roots[0]
	n.Add(1)
	go rapidu.Walk(root, &n, sizes)

	go func() {
		n.Wait()
		close(sizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, total int64

loop:
	for {
		select {
		case size, ok := <-sizes:
			if !ok {
				break loop
			}
			nfiles++
			total += size
		case <-tick:
			fmt.Print(rapidu.PrintDiskUsage(root, nfiles, total))
		}
	}

	fmt.Print(rapidu.PrintDiskUsage(root, nfiles, total))
}

func printUsage() {
	flag.Usage()
	os.Exit(1)
}
