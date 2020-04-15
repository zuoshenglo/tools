package main

import (
	"fmt"
	"github.com/zuoshenglo/tools/file"
)

func main()  {
	//file.DownFile("xxxxxxx", "/tmp/config")
	fmt.Printf(file.GetFileMd5Sum("/tmp/config"))
}
