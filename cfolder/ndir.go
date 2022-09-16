package main

// cdir is a small program that checks and return the numbers of subdirectories in a directory

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type Counter struct {
	count int64
}

func main() {
	l := Counter{}
	val := flag.String("n", "", "pass in a directory name")
	flag.Parse()
	if check(val) {
		fmt.Println("Usage: ndir -n directory_name")
		return
	}

	c := l.CountDir(*(val))
	fmt.Println(c)

}

func (l *Counter) CountDir(path string) int64 {
	var wg sync.WaitGroup

	rd, err := os.ReadDir(path)
	if err != nil {
		CustomError(path, err)
	}

	for _, file := range rd {
		npath := filepath.Join(path, file.Name())
		if file.IsDir() {
			l.count += 1
			wg.Add(1)
			go func() {
				l.CountDir(npath)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	return l.count
}

func check(path *string) bool {
	if *path == "" {
		return true
	}

	return false
}

func CustomError(path string, err error) {
	if os.IsNotExist(err) {
		fmt.Println("No such file or directory")
		return
	}
}
