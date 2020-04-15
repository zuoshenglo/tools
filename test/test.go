package main

import "github.com/zuoshenglo/tools/file"

func main()  {
	file.DownFile("http://service-k8s-key-manager.devops.dev.dm-ai.cn/api/v1/get-k8s-key-file?env=dev", "/tmp/config1111")
	//fmt.Printf(file.GetFileMd5Sum("/tmp/config"))
}
