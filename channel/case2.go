package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var token = make(chan int, 100)

// 入参是一个文件目录，一个INT64的只接收的单向channel
func walkDir(dir string, fileSizes chan<- int64, n *sync.WaitGroup) {
	defer n.Done()
	for _, entry := range dirEntries(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, fileSizes, n)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirEntries(dir string) []os.FileInfo {
	token <- 1
	defer func() { <-token }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

func main() {
	var nfiles, nbytes int64
	var n sync.WaitGroup

	root := ""
	verbose := false
	t1 := time.Now()
	fileSizes := make(chan int64)
	tick := make(<-chan time.Time)

	flag.StringVar(&root, "p", ".", "input dir.")
	flag.BoolVar(&verbose, "v", false, "add verbose if you want")
	flag.Parse()

	if verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	n.Add(1)
	go walkDir(root, fileSizes, &n)

	go func() {
		n.Wait()
		close(fileSizes)
		log.Println("close ...")
	}()

loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			log.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
		}
	}
	log.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
	log.Println(time.Since(t1))
}
