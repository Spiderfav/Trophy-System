package main

import(
  "fmt"
)

type person struct{
  name string
  age int

}
func main(){
  p:= person{name: "Rui", age: 18}
  fmt.Println(p)
  fmt.Println(p.age)

}
