package file

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
)

// down file
func DownFile(url string, localFile string)  {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	f, err := os.Create(localFile)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	io.Copy(f, res.Body)
}

// get file md5 sum

func GetFileMd5Sum(filePath string) string {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Open", err)
		panic(err)
	}

	defer f.Close()
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		fmt.Println("Copy", err)
		panic(err)
	}

	return fmt.Sprintf("%x", md5hash.Sum(nil))
}

// 写入文件内容, 文件不存在创建，存在清0，
func WriteFle(context []byte)  {
	file, _ := os.OpenFile("test2.txt", os.O_RDWR | os.O_TRUNC | os.O_CREATE, 0664)
	defer file.Close()
	count, _ := file.Write(context)
	fmt.Println(count)
	file.Sync()
}
