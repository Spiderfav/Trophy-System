package main

import(
  "fmt"
)

func main(){

  //Declare array
  var a [5]int
  a[2] = 7

  //Declare array and data
  b := [5]int{5,4,3,2,1}

  // Declare "lists"
  c := []int{5,4,3,2,1}

  c = append(c,13)

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
