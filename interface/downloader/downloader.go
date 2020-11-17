package main

import (
	"fmt"
	"go-demo/interface/infra"
)

type retriever interface {
	Get(string) string
}

func getRetriever() retriever {
	return infra.Retriever{}
}

func main() {
	var r retriever = getRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))
}
