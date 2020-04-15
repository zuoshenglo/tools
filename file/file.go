package file

import (
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
	f, err := os.Create(localFile)
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
}

