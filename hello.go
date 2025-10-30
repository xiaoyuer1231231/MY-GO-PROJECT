package main

import "fmt"

func main() {
      var num2 *int
      var num int
      num2=&num
      fmt.Println(num2)
      fmt.Println(num)

}

func method(c int ,b string) ( int,  string) {
   
     return c, b  // 裸返回
}