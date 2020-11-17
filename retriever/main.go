package main

import (
	"fmt"
	"go-demo/retriever/mock"
	real2 "go-demo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("contents: ", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}

func main() {
	var r Retriever
	r = mock.Retriever{Contents: "www.baidu.com"}
	inspect(r)
	r = &real2.Retriever{UserAgent: "Mozilla/5.0", TimeOut: time.Minute}
	inspect(r)

	// type assertion
	realRetriever := r.(*real2.Retriever)
	fmt.Println(realRetriever.TimeOut)

	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	}

	//fmt.Println(download(r))
}
