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

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func post(poster Poster) {
	poster.Post("http://www.imooc.com", map[string]string{"name": "ccmouse", "course": "golang"})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "http://www.imooc.com"

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
	//s.Get()
	//s.Post()
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("contents: ", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}

func main() {
	var r Retriever
	retriever := &mock.Retriever{Contents: "www.baidu.com"}
	r = retriever
	inspect(r)
	r = &real2.Retriever{UserAgent: "Mozilla/5.0", TimeOut: time.Minute}
	inspect(r)

	// type assertion
	realRetriever := r.(*real2.Retriever)
	fmt.Println(realRetriever.TimeOut)

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	}

	fmt.Println("try a session")
	fmt.Println(session(retriever))

	//fmt.Println(download(r))
}
