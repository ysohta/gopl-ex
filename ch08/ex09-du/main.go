package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

type fileInfo struct {
	nfiles, nbytes int64
}

func main() {
	// ...determine roots...

	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	ch := make([]chan int64, len(roots))
	for i, _ := range ch {
		ch[i] = make(chan int64)
	}

	// Traverse each root of the file tree in parallel.
	for i, root := range roots {
		var n sync.WaitGroup
		n.Add(1)
		go walkDir(root, &n, ch[i])
		go func(c chan int64) {
			n.Wait()
			close(c)
		}(ch[i])
	}

	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}

	files := map[string]fileInfo{}
	var wg sync.WaitGroup

	for i, c := range ch {
		wg.Add(1)
		go func(root string, c chan int64) {
			var nfiles, nbytes int64
		loop:
			for {
				select {
				case size, ok := <-c:
					if !ok {
						wg.Done()
						break loop // fileSizes was closed
					}
					nfiles++
					nbytes += size

					f := files[root]
					f.nfiles = nfiles
					f.nbytes = nbytes
					files[root] = f

				case <-tick:
					// printDiskUsage(root, nfiles, nbytes)
					printDiskUsageAll(files)
				}
			}

		}(roots[i], c)

	}

	wg.Wait()
	fmt.Println("-----")

	printDiskUsageAll(files)
}

func printDiskUsageAll(files map[string]fileInfo) {
	for root, f := range files {
		printDiskUsage(root, f.nfiles, f.nbytes)
	}
	fmt.Println()
}

func printDiskUsage(root string, nfiles, nbytes int64) {
	fmt.Printf("%d files\t%.1f GB\t%s\n", nfiles, float64(nbytes)/1e9, root)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
