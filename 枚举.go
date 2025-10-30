package main

import "fmt"

// type Gender  string

// const (test  Gender  ="test1"
// 	testOne  Gender  ="testOne1")
	

const pre int = 1
const a int = iota
const (
    b int = iota
    c
    d
    e
)
const (
    f = 2
    g = iota
    h = iota
    i
)
func main() {
	    fmt.Println("Male:", a)

    fmt.Println("Male:", g)
	 fmt.Println("Male:", h)
}