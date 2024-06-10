package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int      { return len(p) }
func (p people) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// 字典序
func (p people) Less(i, j int) bool { return p[i] < p[j] }

// 字符串长度排序
//func (p people) Less(i, j int) bool { return len(p[i]) < len(p[j]) }

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

}

// https://golang.org/pkg/sort/#Sort
// https://golang.org/pkg/sort/#Interface
