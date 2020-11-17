package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println(m, m2, m3)

	// 遍历
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 取值
	courseName := m["course"]
	fmt.Println(courseName)

	if causeName, ok := m["couse"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key does not exist")
	}

	// 删除
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")

	name, ok = m["name"]
	fmt.Println(name, ok)
}
