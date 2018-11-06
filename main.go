package main

import "fmt"

func main() {
  var age = 23
  //var name = "Bharath"
  // shorthand
//   name := "Bharath"
  size := 1.3
  //var sizes float32 = 2.3
  var isCool = true
  isCool = false
  
   
  name,email := "Bharath", "bharathchandra@gmal.com"

  fmt.Println(name,age,email)

  //for getting the type
  fmt.Printf("%T\n", name)
  fmt.Printf("%T\n", age)
  fmt.Printf("%T\n", isCool)
  fmt.Printf("%T\n", size)
}